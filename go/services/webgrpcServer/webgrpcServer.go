package webgrpcServer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/SOCOMD/ts3Bot"

	grpcMembers "github.com/SOCOMD/staff"
	usr "github.com/SOCOMD/staff/go/db/user"
	"github.com/SOCOMD/staff/go/services/dbClient"
	"github.com/SOCOMD/staff/go/services/ts3BotClient"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	grpcServer *grpc.Server
	httpServer *http.Server
)

func Serve() {
	go serveInternal()
}

func Disconnect() {
	if grpcServer != nil {
		grpcServer.GracefulStop()
	}

	if httpServer != nil {
		httpServer.Close()
	}
}

func serveInternal() {
	webgrpcHost := os.Getenv("HOST")
	webgrpcHostOverride := os.Getenv("HOST")
	if len(webgrpcHostOverride) > 0 {
		fmt.Printf("Using webgrpc host override!")
		webgrpcHost = webgrpcHostOverride
	}

	webgrpcAddr := fmt.Sprintf("%s:%s", webgrpcHost, os.Getenv("PORT_WEBGRPC"))

	grpcServer = grpc.NewServer()
	grpcMembers.RegisterMembersServer(grpcServer, &memberService{})

	grpclog.SetLogger(log.New(os.Stdout, "webgrpcServer: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    webgrpcAddr,
		Handler: http.HandlerFunc(handler),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("Failed to start http server: %v", err)
	}
}

type memberService struct{}

func (s *memberService) GetUser(ctx context.Context, userQuery *grpcMembers.GetUserMessage) (user *grpcMembers.User, err error) {

	dbInstance, dberr := dbClient.GetDBInstance()
	if dberr != nil {
		return nil, dberr
	}

	ts3BotInstance, ts3err := ts3BotClient.GetTs3BotInstance()
	if ts3err != nil {
		return nil, ts3err
	}

	id, _ := strconv.Atoi(userQuery.Id)
	dbUser, err := usr.Get(id, dbInstance)
	if err != nil {
		return nil, err
	}

	var ts3Query ts3Bot.User
	ts3Query.Uuid = *dbUser.TeamspeakUUID

	ts3User, ts3err := ts3BotInstance.GetUser(context.Background(), &ts3Query)
	if ts3err != nil {
		return nil, err
	}

	var retUsr grpcMembers.User
	retUsr.TsName = ts3User.Name
	retUsr.Tsuuid = *dbUser.TeamspeakUUID
	retUsr.Email = *dbUser.Email

	return &retUsr, nil
}
