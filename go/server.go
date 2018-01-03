package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/SOCOMD/staff"
	dbUser "github.com/SOCOMD/staff/go/db/user"
	"github.com/SOCOMD/ts3Bot"
	"github.com/jmoiron/sqlx"
)

type server struct {
	db           *sqlx.DB
	ts3botClient ts3Bot.Ts3BotClient
}

// Wrapper Implementations

func (s *server) SetDatabase(db *sql.DB) {
	sqlx.NewDb(db, "mysql")
}

func (s server) GetDatabase() *sql.DB {
	if s.db == nil {
		return nil
	}
	return s.db.DB
}

func (s *server) SetTSBot(client ts3Bot.Ts3BotClient) {
	s.ts3botClient = client
}

// Service Implimentations
func (s *server) GetUser(ctx context.Context, userQuery *staff.GetUserRequest) (res *staff.User, err error) {
	res = &staff.User{}
	// validate token.
	steamid, err := validateToken(userQuery.Token)
	if err != nil {
		return
	}
	var DBUser dbUser.User
	// get search type.
	switch userQuery.Type {
	case staff.GetUserRequest_ID:
		DBUser, err = dbUser.Get(dbUser.FieldID, userQuery.Search, s.db)
	case staff.GetUserRequest_STEAMID:
		DBUser, err = dbUser.Get(dbUser.FieldSteamID, userQuery.Search, s.db)
	case staff.GetUserRequest_TSDBID:
		DBUser, err = dbUser.Get(dbUser.FieldTSDBID, userQuery.Search, s.db)
	case staff.GetUserRequest_TSUUID:
		DBUser, err = dbUser.Get(dbUser.FieldTSUUID, userQuery.Search, s.db)
	case staff.GetUserRequest_TOKEN:
		DBUser, err = dbUser.Get(dbUser.FieldSteamID, steamid, s.db)
	default:
		err = fmt.Errorf("Invalid Search Type")
	}
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	res.Id = strconv.Itoa(DBUser.ID)
	res.Email = *DBUser.Email
	res.Active = DBUser.Active > 0
	res.Admin = DBUser.Admin
	res.Dob = *DBUser.DoB
	res.Gender = *DBUser.Gender
	res.Joindate = *DBUser.JoinDate
	res.Steamid = *DBUser.SteamID
	// it is possible we do not have a teamspeak user assosiated with this person.
	if len(*DBUser.TeamspeakID) == 0 && len(*DBUser.TeamspeakUUID) == 0 {
		return
	}
	tsUser := &ts3Bot.User{}
	tsUser.Uuid = *DBUser.TeamspeakUUID
	tsUser.Dbid = *DBUser.TeamspeakID

	tsUser, ts3err := s.ts3botClient.GetUser(ctx, tsUser)
	if ts3err != nil {
		return
	}
	res.TsName = tsUser.Name
	res.Tsdbid = tsUser.Dbid
	res.Tsuuid = tsUser.Uuid
	res.Tscreated = tsUser.Created
	res.Tslastconnected = tsUser.Lastconnected

	return
}

func (s *server) AuthStatus(ctx context.Context, in *staff.GetAuthStatusRequest) (res *staff.GetAuthStatusResult, err error) {
	return nil, fmt.Errorf("Not Implemented")
}
