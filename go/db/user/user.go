package user

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/SOCOMD/ts3Bot"
	"github.com/jmoiron/sqlx"
)

//User object for db query
type User struct {
	ID               int     `db:"id"`
	SteamID          *string `db:"steamid"`
	TeamspeakID      *string `db:"tsdbid"`
	TeamspeakUUID    *string `db:"tsuuid"`
	Email            *string `db:"email"`
	Password         *string `db:"password"`
	JoinDate         *string `db:"joindate"`
	DoB              *string `db:"dob"`
	Gender           *string `db:"gender"`
	Admin            *int    `db:"admin"`
	Active           int     `db:"active"`
	AttendenceCredit int     `db:"attendenceCredit"`
}

var (
	emptyString string = ""
)

func (u *User) SafeString(value *string) (result string) {
	if value == nil {
		return ""
	}
	return *value
}

//GetAll returns a list of all users from the database
func GetAll(db *sql.DB) (result []User, err error) {
	dbx := sqlx.NewDb(db, "mysql")
	err = dbx.Select(&result, "SELECT * FROM user")
	for _, user := range result {
		user.reflect()
	}
	return
}

func (u *User) reflect() {
	val := reflect.ValueOf(u).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		if valueField.Kind() == reflect.Ptr {
			if valueField.IsNil() {
				valueField.Set(reflect.New(typeField.Type.Elem()))
			}
		}
	}
}

//Get returns a single user from the database
func Get(id int, db *sql.DB) (result User, err error) {
	dbx := sqlx.NewDb(db, "mysql")
	err = dbx.Get(&result, "SELECT * FROM user WHERE id=?", id)
	result.reflect()
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
	return
}

//ConvertUserSteamToDB converts a steam id to a db user
func ConvertUserSteamToDB(steamid string, db *sql.DB) (user User, err error) {
	dbx := sqlx.NewDb(db, "mysql")
	err = dbx.Get(&user, "SELECT * FROM user WHERE steamid=?", steamid)
	return
}

//Validate validate user by steamid and ts3uuid
func Validate(steamid string, ts3Uuid string, db *sql.DB) (user User, err error) {
	dbx := sqlx.NewDb(db, "mysql")
	err = dbx.Get(&user, "SELECT * FROM user WHERE steamid=? OR tsuuid=?", steamid, ts3Uuid)
	return
}

//Register a user into the DB with steamid and ts3uuid
func Register(steamid string, ts3Uuid string, ts3Client ts3Bot.Ts3BotClient, db *sql.DB) (err error) {

	var ts3QueryUser ts3Bot.User
	ts3QueryUser.Uuid = ts3Uuid

	ts3User, ts3Err := ts3Client.GetUser(context.Background(), &ts3QueryUser)
	if ts3Err != nil {
		return ts3Err
	}

	dbx := sqlx.NewDb(db, "mysql")
	_, err = dbx.Exec("INSERT INTO user (steamid,tsdbid,tsuuid) VALUES (?,?,?)", steamid, ts3User.Dbid, ts3User.Uuid)
	return
}
