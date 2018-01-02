package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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
	testcase()
	cleanup()

	//hostWebsite()
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

func initialise() {
	//Establish DB Connection
	envDBHost := os.Getenv("DBHOST")
	envDBUser := os.Getenv("DBUSER")
	envDBPass := os.Getenv("DBPASS")
	envDBName := os.Getenv("DBNAME")
	envTs3BotAddr := os.Getenv("TS3BOTADDR")

	var err error
	db, err = sql.Open("mysql", envDBUser+":"+envDBPass+"@tcp("+envDBHost+":3306)/"+envDBName)
	if err != nil {
		panic(err.Error())
	}

	//Establish ts3BotClient Connection
	ts3grpc, err = grpc.Dial(envTs3BotAddr, grpc.WithInsecure())
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

func testcase() {

	//// Get Sinle TS3 User by setting search user data
	// var ts3SearchUser ts3Bot.User
	// ts3SearchUser.Uuid = `vNfQp42gNSmnZXzwycVTamagAvE=`
	// ts3User, ts3Err := ts3Client.GetUser(context.Background(), &ts3SearchUser)
	// if ts3Err != nil {
	// 	fmt.Println(ts3Err.Error())
	// 	return
	// }

	// fmt.Println("Entry:", " TSNAME:", ts3User.Name, " TSDBID:", ts3User.Dbid, " TSUUID:", ts3User.Uuid, " TSLASTCONN:", ts3User.Lastconnected)

	////Get all ts3clients and add them to the DB
	// ts3UserList, ts3Err := ts3Client.GetUsers(context.Background(), &ts3Bot.Nil{})
	// if ts3Err != nil {
	// 	fmt.Println(ts3Err.Error())
	// 	return
	// }

	// for _, ts3User := range ts3UserList.Users {
	// 	dbUser, err := user.ConvertUserTs3ToDB(ts3User, db)
	// 	if err != nil {
	// 		user.AddTs3User(ts3User, db)
	// 		dbUser, err = user.ConvertUserTs3ToDB(ts3User, db)
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 			continue
	// 		}
	// 	}

	// 	fmt.Println("Entry:", " ID:", dbUser.ID, " TSID:", dbUser.SafeString(dbUser.TeamspeakID), " TSNAME:", ts3User.Name, " STEAMID:", dbUser.SafeString(dbUser.SteamID))

	// }
}
