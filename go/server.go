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
	res = &staff.GetAuthStatusResult{}
	steamid, err := validateToken(in.GetToken())
	u, err := dbUser.Get(dbUser.FieldSteamID, steamid, s.db)
	if err != nil {
		return
	}
	res.Admin = u.Admin
	return
}

func (s *server) UpdateUser(ctx context.Context, userQuery *staff.UpdateUserRequest) (res *staff.NilResult, err error) {
	fmt.Println("Update User Called")
	res = &staff.NilResult{}

	if userQuery.User == nil {
		return nil, fmt.Errorf("User was nil")
	}

	steamid, err := validateToken(userQuery.Token)
	if err != nil {
		return
	}

	DBUser, err := dbUser.Get(dbUser.FieldSteamID, steamid, s.db)
	if err != nil {
		return
	}

	webUser := userQuery.User

	if webUser.Tsuuid != "" {
		var ts3Query ts3Bot.User
		ts3Query.Uuid = webUser.Tsuuid
		ts3User, ts3err := s.ts3botClient.GetUser(context.Background(), &ts3Query)
		if ts3err != nil {
			return
		}

		*DBUser.TeamspeakID = ts3User.Dbid
		*DBUser.TeamspeakUUID = webUser.Tsuuid
	} else {
		DBUser.TeamspeakID = nil
		DBUser.TeamspeakUUID = nil
	}

	fmt.Printf("%#v\n", userQuery.User)

	*DBUser.Email = webUser.Email
	*DBUser.JoinDate = webUser.Joindate
	*DBUser.DoB = webUser.Dob
	*DBUser.Gender = webUser.Gender
	if webUser.Active == true {
		DBUser.Active = 1
	} else {
		DBUser.Active = 0
	}

	DBUser.Admin = webUser.Admin

	err = dbUser.Update(DBUser, s.db)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
