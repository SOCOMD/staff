package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/SOCOMD/staff/go/db/user"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func main() {

	dbTest()

	http.HandleFunc("/", handler)
	log.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("../website/dist/"))
	if strings.Contains(r.URL.String(), ".") == false {
		r.URL.Path = "/"
	}
	fmt.Println(r.URL.String())
	fs.ServeHTTP(w, r)
}

func dbTest() {

	fmt.Println("Running DB Tests")

	dbHost := os.Getenv("DBHOST")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")

	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":3306)/"+dbName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	u, err := user.GetSingle(1, db)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("User credits =", u.AttendenceCredit)
	}

	u.AttendenceCredit++

	err = user.UpdateSingle(u, db)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Updated User credits =", u.AttendenceCredit)
	}
}
