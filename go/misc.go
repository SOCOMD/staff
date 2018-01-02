package main

import (
	"context"
	"fmt"

	"github.com/SOCOMD/staff/go/db/user"
	"github.com/SOCOMD/ts3Bot"
)

func testcase() {
	ts3UserList, ts3Err := ts3Client.GetUsers(context.Background(), &ts3Bot.Nil{})
	if ts3Err != nil {
		fmt.Println(ts3Err.Error())
		return
	}

	for _, ts3User := range ts3UserList.Users {
		dbUser, err := user.ConvertUserTs3ToDB(ts3User, db)
		if err != nil {
			user.AddTs3User(ts3User, db)
			dbUser, err = user.ConvertUserTs3ToDB(ts3User, db)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		}

		fmt.Println("Entry:", " ID:", dbUser.ID, " TSID:", dbUser.SafeString(dbUser.TeamspeakID), " TSNAME:", ts3User.Name, " STEAMID:", dbUser.SafeString(dbUser.SteamID))

	}
}
