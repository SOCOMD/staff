package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/SOCOMD/env"

	"github.com/SOCOMD/staff/go/services/dbClient"
	"github.com/SOCOMD/staff/go/services/ts3BotClient"
	"github.com/SOCOMD/staff/go/services/webgrpcServer"
	_ "github.com/go-sql-driver/mysql"
)

var (
	e          env.Env
	webAddress string
)

func main() {
	helpFlag := flag.Bool("help", false, "If Defined will print the help menu")
	flag.Parse()
	e = env.Get()
	if *helpFlag == true {
		//print all help things and leave
		fmt.Println(e)
		return
	}
	webAddress = e.Staff.WebHost
	if len(e.Staff.WebPort) > 0 {
		webAddress += ":" + e.Staff.WebPort
	}
	initialise()
	cleanup()
}

func initialise() {
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

	log.Printf("Listening on %s:%s", e.Staff.WebHost, e.Staff.WebPort)
	log.Println(http.ListenAndServe(e.Staff.WebHost+":"+e.Staff.WebPort, nil))
}
