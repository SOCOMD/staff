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
func (s *server) GetUser(ctx context.Context, userQuery *staff.GetUserMessage) (*staff.User, error) {

	id, _ := strconv.Atoi(userQuery.Id)
	DbUser, err := dbUser.Get(id, s.db.DB)
	if err != nil {
		return nil, fmt.Errorf("No User Found")
	}

	var tsUser *ts3Bot.User
	tsUser.Uuid = *DbUser.TeamspeakUUID

	tsUser, ts3err := s.ts3botClient.GetUser(ctx, tsUser)
	if ts3err != nil {
		return nil, err
	}

	var retUsr staff.User
	retUsr.Id = strconv.Itoa(DbUser.ID)
	retUsr.TsName = tsUser.Name
	retUsr.Tsdbid = tsUser.Dbid
	retUsr.Tsuuid = tsUser.Uuid
	retUsr.Tscreated = tsUser.Created
	retUsr.Tslastconnected = tsUser.Lastconnected
	retUsr.Email = *DbUser.Email
	retUsr.Active = DbUser.Active > 0
	retUsr.Admin = DbUser.Admin
	retUsr.Dob = *DbUser.DoB
	retUsr.Gender = *DbUser.Gender
	retUsr.Joindate = *DbUser.JoinDate

	return &retUsr, nil
}
