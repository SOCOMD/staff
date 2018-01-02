package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/SOCOMD/ts3Bot"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

var (
	db         *sql.DB
	ts3grpc    *grpc.ClientConn
	ts3Client  ts3Bot.Ts3BotClient
	webAddress string
	jwtsecret  []byte
)

func main() {
	initialise()
	hostWebsite()
	//testcase()
	cleanup()

}

func hostWebsite() {
	http.HandleFunc("/steamlogin", steamLoginHandler)
	http.HandleFunc("/steamcallback", steamCallbackHandler)
	http.HandleFunc("/", handler)
	log.Println("Listening on", webAddress)
	log.Println(http.ListenAndServe(webAddress, nil))
}

func initialise() {
	//Establish DB Connection
	envDBHost := os.Getenv("DBHOST")
	envDBUser := os.Getenv("DBUSER")
	envDBPass := os.Getenv("DBPASS")
	envDBName := os.Getenv("DBNAME")
	envTs3BotAddr := os.Getenv("TS3BOTADDR")
	webAddress = os.Getenv("MEMBERS_WEBHOST") + ":" + os.Getenv("MEMBERS_WEBPORT")
	jwtsecret = []byte(os.Getenv("MEMBERS_JWTSECRET"))

	var err error
	db, err = sql.Open("mysql", envDBUser+":"+envDBPass+"@tcp("+envDBHost+":3306)/"+envDBName)
	if err != nil {
		log.Println(err.Error())
	}

	//Establish ts3BotClient Connection
	ts3grpc, err = grpc.Dial(envTs3BotAddr, grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}

	ts3Client = ts3Bot.NewTs3BotClient(ts3grpc)
	if ts3Client == nil {
		log.Println(err.Error())
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
