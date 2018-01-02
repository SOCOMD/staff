package dbClient

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	dbInstance *sql.DB
)

func Serve() {
	dbHost := os.Getenv("HOST")
	dbHostOverride := os.Getenv("HOST_DB")
	if len(dbHostOverride) > 0 {
		fmt.Printf("Using db host override!")
		dbHost = dbHostOverride
	}

	dbPort := os.Getenv("PORT_DB")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	mysqlAddr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	dbInstance, err = sql.Open("mysql", mysqlAddr)
	if err != nil {
		log.Println(err.Error())
	}
}

func Disconnect() {
	if dbInstance != nil {
		dbInstance.Close()
	}
}

func GetDBInstance() (db *sql.DB) {
	return dbInstance
}
