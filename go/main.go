package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/SOCOMD/staff/go/db/user"
	"github.com/SOCOMD/ts3Bot"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

var (
	db        *sql.DB
	ts3grpc   *grpc.ClientConn
	ts3Client ts3Bot.Ts3BotClient
)

func main() {
	initialise()
	//generateDBEntries()
	getDBClientInfo(11)
	cleanup()

	hostWebsite()
}

func initialise() {
	//Establish DB Connection
	dbHost := os.Getenv("DBHOST")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")
	ts3BotAddr := os.Getenv("TS3BOTADDR")

	var err error
	db, err = sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+":3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	//Establish ts3BotClient Connection
	ts3grpc, err = grpc.Dial(ts3BotAddr, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	ts3Client = ts3Bot.NewTs3BotClient(ts3grpc)
	if ts3Client == nil {
		panic(err.Error())
	}
}

func cleanup() {
	if db != nil {
		db.Close()
	}

	if ts3grpc != nil {
		ts3grpc.Close()
	}
}

func hostWebsite() {
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

func generateDBEntries() {
	if ts3Client == nil {
		panic("invalud ts3client")
	}

	userList, err := ts3Client.GetUsers(context.Background(), &ts3Bot.Nil{})
	if err != nil {
		panic(err.Error())
	}

	for _, ts3User := range userList.Users {
		exists := user.Exists_TS3(ts3User.Dbid, db)
		if exists == false {
			user.Register_TS3(ts3User.Dbid, db)
			continue
		}
	}

	dbUsers, err := user.GetAll(db)
	if err != nil {
		panic(err.Error())
	}

	for _, dbUser := range dbUsers {
		var tsUser ts3Bot.User
		tsUser.Dbid = dbUser.TeamspeakID

		tsUserPtr, err := ts3Client.GetUser(context.Background(), &tsUser)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println("User:", " DB_ID:", dbUser.ID, " TS_ID:", dbUser.TeamspeakID, " TS_USERNAME:", tsUserPtr.Name)
	}
}

func getDBClientInfo(dbid int) {
	dbUser, err := user.GetSingle(dbid, db)
	if err != nil {
		panic(err.Error())
	}

	var tsUser ts3Bot.User
	tsUser.Dbid = dbUser.TeamspeakID
	tsUserPtr, err := ts3Client.GetUser(context.Background(), &tsUser)
	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Println("User:", " DB_ID:", dbUser.ID, " TS_ID:", dbUser.TeamspeakID, " TS_USERNAME:", tsUserPtr.Name)
}
