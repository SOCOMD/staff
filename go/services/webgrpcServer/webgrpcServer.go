package webgrpcServer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"google.golang.org/grpc/codes"

	grpcMembers "github.com/SOCOMD/staff"
	dbUser "github.com/SOCOMD/staff/go/db/user"
	"github.com/SOCOMD/staff/go/services/dbClient"
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

	dbInstance := dbClient.GetDBInstance()
	if dbInstance == nil {
		return nil, grpc.Errorf(codes.Unavailable, "MySQL Database Instance is nil")
	}

	id, _ := strconv.Atoi(userQuery.Id)
	res, err := dbUser.Get(id, dbInstance)
	if err != nil {
		return nil, err
	}

	var retUsr grpcMembers.User
	retUsr.TsName = *res.TeamspeakID
	retUsr.Tsuuid = *res.TeamspeakUUID
	retUsr.Email = *res.Email

	return &retUsr, nil
}
