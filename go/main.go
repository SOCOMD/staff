package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SOCOMD/staff/go/services/dbClient"
	"github.com/SOCOMD/staff/go/services/ts3BotClient"
	"github.com/SOCOMD/staff/go/services/webgrpcServer"
	_ "github.com/go-sql-driver/mysql"
)

var (
	webAddress string
	jwtsecret  []byte
)

func main() {
	initialise()
	cleanup()
}

func initialise() {

	webHost := os.Getenv("HOST")
	webHostOverride := os.Getenv("HOST_WEBHOST")
	if len(webHostOverride) > 0 {
		fmt.Printf("Using WebHost override!")
		webHost = webHostOverride
	}

	webAddress = fmt.Sprintf("%s:%s", webHost, os.Getenv("PORT_WEBHOST"))
	jwtsecret = []byte(os.Getenv("ACC_JWTSECRET"))

	//Connect Clients
	dbClient.Connect()
	ts3BotClient.Connect()

	//Host Servers
	webgrpcServer.Serve()
	hostWebsite()

}

func cleanup() {
	//Disconnect Clients
	dbClient.Disconnect()
	ts3BotClient.Disconnect()

	//Disconnect Servers
	webgrpcServer.Disconnect()
}

func hostWebsite() {
	http.HandleFunc("/steamlogin", steamLoginHandler)
	http.HandleFunc("/steamcallback", steamCallbackHandler)
	http.HandleFunc("/", handler)

	log.Println("Listening on", webAddress)
	log.Println(http.ListenAndServe(webAddress, nil))
}
