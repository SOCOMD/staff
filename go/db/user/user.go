package user

import (
	"database/sql"
	"fmt"

	"github.com/SOCOMD/ts3Bot"
	"github.com/jmoiron/sqlx"
)

//User object for db query
type User struct {
	ID               int     `db:"id"`
	SteamID          *string `db:"steamid"`
	TeamspeakID      *string `db:"tsdbid"`
	Email            *string `db:"email"`
	Password         *string `db:"password"`
	JoinDate         *string `db:"joindate"`
	DoB              *string `db:"dob"`
	Gender           *string `db:"gender"`
	Admin            int     `db:"admin"`
	Active           int     `db:"active"`
	AttendenceCredit int     `db:"attendenceCredit"`
}

func (u *User) SafeString(value *string) (result string) {
	const nilReturn = "isNil"
	if value == nil {
		result = nilReturn
		return
	}
	result = *value
	return
}

//GetAll returns a list of all users from the database
func GetAll(db *sql.DB) (result []User, err error) {
	dbx := sqlx.NewDb(db, "mysql")
	err = dbx.Select(&result, "SELECT * FROM user")
	return
}

//Get returns a single user from the database
func Get(id int, db *sql.DB) (result User, err error) {
	dbx := sqlx.NewDb(db, "mysql")
	err = dbx.Get(&result, "SELECT * FROM user WHERE id=?", id)
	return
}

//UpdateAll will update multiple users at once in the db
func UpdateAll(users []User, db *sql.DB) {
	for _, user := range users {
		updateErr := Update(user, db)
		if updateErr != nil {
			fmt.Println(updateErr.Error())
		}
	}
}

//Update updates a single user in the db
func Update(user User, db *sql.DB) (err error) {
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

//ConvertUserTs3ToDB converts a TS3 user to a db user
func ConvertUserTs3ToDB(ts3User *ts3Bot.User, db *sql.DB) (user User, err error) {
	if ts3User == nil {
		err = fmt.Errorf("ts3 user is nil")
		return
	}

	dbx := sqlx.NewDb(db, "mysql")
	err = dbx.Get(&user, "SELECT * FROM user WHERE tsdbid=?", ts3User.Dbid)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

//ConvertUserSteamToDB converts a steam id to a db user
func ConvertUserSteamToDB(steamid string, db *sql.DB) (user User, err error) {
	dbx := sqlx.NewDb(db, "mysql")
	err = dbx.Get(&user, "SELECT * FROM user WHERE steamid=?", steamid)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
