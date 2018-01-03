package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/SOCOMD/ts3Bot"
	"google.golang.org/grpc"
)

type grpcServerWrapper interface {
	SetDatabase(*sql.DB)
	GetDatabase() *sql.DB
	SetTSBot(ts3Bot.Ts3BotClient)
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//var wrapper grpcServerWrapper
	server, ok := info.Server.(*server)
	//wrapper = server
	if ok == false {
		err = fmt.Errorf("Requested Service does not implement wrapper interface: %#v", server)
		return
	}
	// check database connection is still good,
	// if not replace it before doing the call
	if server.db == nil {
		db, err := openDatabase()
		if err != nil {

			return nil, err
		}
		server.db = db
	}

	// ping failed, create new conn, then test and set if valid or close
	if server.db.Ping() != nil {
		db, err := openDatabase()
		if err != nil {
			return nil, err
		}
		err = db.Ping()
		if err != nil {
			return nil, err
		}
		server.db = db
	}

	// set a new TSClient Connection every request, to avoid blocking on a single pipe.
	tsbotClient, grpcConn, err := makeTSClient()
	if err != nil {
		return
	}
	server.ts3botClient = tsbotClient
	resp, err = handler(ctx, req)
	grpcConn.Close()
	return
}

func openDatabase() (*sqlx.DB, error) {
	dbinfo := e.Staff.Database
	return sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbinfo.User, dbinfo.Pass, dbinfo.Host, dbinfo.Port, dbinfo.Name))
}

func makeTSClient() (client ts3Bot.Ts3BotClient, grpcConn *grpc.ClientConn, err error) {
	grpcConn, err = grpc.Dial(e.Tsbot.GrpcHost+":"+e.Tsbot.GrpcPort, grpc.WithInsecure())
	if err != nil {
		return
	}
	client = ts3Bot.NewTs3BotClient(grpcConn)
	return
}
