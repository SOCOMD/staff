package user

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/SOCOMD/ts3Bot"
	"github.com/jmoiron/sqlx"
)

const (

	// Field Const's denote database fields

	FieldID       field = "id"
	FieldSteamID  field = "steamid"
	FieldTSDBID   field = "tsdbid"
	FieldTSUUID   field = "tsuuid"
	FieldEmail    field = "email"
	FieldPassword field = "password"
	FieldJoindate field = "joindate"
	FieldDOB      field = "dob"
	FieldGender   field = "gender"
	FieldAdmin    field = "admin"
	FieldActive   field = "active"
)

var (
	emptyString string
)

type field string

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
	Admin            int32   `db:"admin"`
	Active           int32   `db:"active"`
	AttendenceCredit *int32  `db:"attendenceCredit"`
}

func (u *User) SafeString(value *string) (result string) {
	if value == nil {
		return ""
	}
	return *value
}

//GetAll returns a list of all users from the database
func GetAll(db *sqlx.DB) (result []User, err error) {
	err = db.Select(&result, "SELECT * FROM user")
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
func Get(Type field, value string, db *sqlx.DB) (result User, err error) {
	err = db.Get(&result, fmt.Sprintf("SELECT * FROM user WHERE %s=?", Type), value)
	result.reflect()
	return
}

//Update updates a single user in the db
func Update(user User, db *sqlx.DB) (err error) {
	_, err = db.Exec(`UPDATE user SET
		tsuuid=?,
		tsdbid=?,
		email=?,
		joindate=?,
		dob=?,
		gender=?,
		admin=?,
		active=?,
		attendenceCredit=?
		WHERE steamid=?`,
		user.TeamspeakUUID,
		user.TeamspeakID,
		user.Email,
		user.JoinDate,
		user.DoB,
		user.Gender,
		user.Admin,
		user.Active,
		user.AttendenceCredit,
		user.SteamID)
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
func Register(steamid string, db *sqlx.DB) (err error) {
	_, err = db.Exec("INSERT INTO user (steamid) VALUES (?)", steamid)
	return
}
