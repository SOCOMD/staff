package user

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

//User object for db query
type User struct {
	ID               int    `db:"id"`
	Steamid          string `db:"steamid"`
	Tsdbid           string `db:"tsdbid"`
	Email            string `db:"email"`
	Joindate         string `db:"joindate"`
	Dob              string `db:"dob"`
	Gender           string `db:"gender"`
	Admin            int    `db:"admin"`
	AttendenceCredit int    `db:"attendenceCredit"`
	Password         string `db:"password"`
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
