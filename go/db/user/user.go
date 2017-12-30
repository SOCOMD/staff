package user

import (
	"database/sql"
	"fmt"

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
		joindate=?,
		dob=?,
		gender=?,
		admin=?,
		attendenceCredit=?,
		password=?
		WHERE id=?`,
		user.Steamid,
		user.Tsdbid,
		user.Email,
		user.Joindate,
		user.Dob,
		user.Gender,
		user.Admin,
		user.AttendenceCredit,
		user.Password,
		user.ID)
	return
}
