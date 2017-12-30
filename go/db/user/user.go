package user

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

//User object for db query
type User struct {
	ID               int    `db:"id"`
	SteamID          string `db:"steamid"`
	TeamspeakID      string `db:"tsdbid"`
	Email            string `db:"email"`
	Password         string `db:"password"`
	JoinDate         string `db:"joindate"`
	DoB              string `db:"dob"`
	Gender           string `db:"gender"`
	Admin            int    `db:"admin"`
	Active           int    `db:"active"`
	AttendenceCredit int    `db:"attendenceCredit"`
}

func Exists_TS3(tsdbid string, db *sql.DB) (result bool) {
	result = false

	dbx := sqlx.NewDb(db, "mysql")
	err := dbx.Get(&result, "SELECT * FROM user WHERE tsdbid=?", tsdbid)
	if err == nil {
		result = true
		return
	}

	return
}

//GetAll returns a list of all users from the database
func GetAll(db *sql.DB) (result []User, err error) {
	dbx := sqlx.NewDb(db, "mysql")
	err = dbx.Select(&result, "SELECT * FROM user")
	return
}

//GetSingle returns a single user from the database
func GetSingle(id int, db *sql.DB) (result User, err error) {
	dbx := sqlx.NewDb(db, "mysql")
	err = dbx.Get(&result, "SELECT * FROM user WHERE id=?", id)
	return
}

//UpdateMultiple will update multiple users at once in the db
func UpdateMultiple(users []User, db *sql.DB) {
	for _, user := range users {
		updateErr := UpdateSingle(user, db)
		if updateErr != nil {
			fmt.Println(updateErr.Error())
		}
	}
}

//UpdateSingle updates a single user in the db
func UpdateSingle(user User, db *sql.DB) (err error) {
	dbx := sqlx.NewDb(db, "mysql")
	_, err = dbx.Exec(`UPDATE user SET 
		steamid=?,
		tsdbid=?,
		email=?,
		password=?,
		joindate=?,
		dob=?,
		gender=?,
		admin=?,
		active=?,
		attendenceCredit=?
		WHERE id=?`,
		user.SteamID,
		user.TeamspeakID,
		user.Email,
		user.Password,
		user.JoinDate,
		user.DoB,
		user.Gender,
		user.Admin,
		user.Active,
		user.AttendenceCredit,
		user.ID)
	return
}

//Register_TS3 binds a tsdbid to a db id
func Register_TS3(tsdbid string, db *sql.DB) (err error) {
	dbx := sqlx.NewDb(db, "mysql")
	_, err = dbx.Exec(`INSERT INTO user (tsdbid) VALUES(?)`, tsdbid)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
