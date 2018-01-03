package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/SOCOMD/env"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	pb "github.com/SOCOMD/staff"
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

	// create grpc server
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	// register our service onto the grpc server. (can register many services if needed)
	pb.RegisterStaffServer(grpcServer, &server{})

	// wrap our grpc server in grpc-web so the frontend can talk to it.
	wrappedServer := grpcweb.WrapServer(grpcServer)

	// setup our http server to have a custom handler so each request can first be checked if its
	// a grpc request before its passed off to the default http serve mux.
	httpServer := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// if it is a grpc matching request handle it,
			// otherwise treat as a standard http request
			if wrappedServer.IsGrpcWebRequest(r) {
				wrappedServer.ServeHTTP(w, r)
				return
			}
			http.DefaultServeMux.ServeHTTP(w, r)
		}),
		Addr: webAddress,
	}
	// register http handlers
	http.HandleFunc("/steamlogin", steamLoginHandler)
	http.HandleFunc("/steamcallback", steamCallbackHandler)
	http.HandleFunc("/", handler)
	log.Printf("Listening on %s", webAddress)
	log.Println(httpServer.ListenAndServe())
}
