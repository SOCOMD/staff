package main

import (
	dbUser "github.com/SOCOMD/staff/go/db/user"
)

func testcase() {

}

func registerUser(steamid string) (err error) {
	db, err := openDatabase()
	if err != nil {
		return err
	}
	_, err = dbUser.Get(dbUser.FieldSteamID, steamid, db)
	if err == nil {
		return
	}
	err = dbUser.Register(steamid, db)
	return
}
