package ts3BotClient

import (
	"fmt"
	"os"

	"github.com/SOCOMD/ts3Bot"
	"google.golang.org/grpc"
)

var (
	grpcClient *grpc.ClientConn
	ts3Client  ts3Bot.Ts3BotClient
)

func Connect() (err error) {
	//Establish DB Connection
	ts3BotHost := os.Getenv("HOST")
	ts3BotHostOverride := os.Getenv("HOST_TS3BOT")
	if len(ts3BotHostOverride) > 0 {
		fmt.Printf("Using ts3 bot host override!")
		ts3BotHost = ts3BotHostOverride
	}

	//Establish ts3BotClient Connection
	ts3BotAddr := fmt.Sprintf("%s:%s", ts3BotHost, os.Getenv("PORT_TS3BOT"))
	grpcClient, err := grpc.Dial(ts3BotAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	ts3Client = ts3Bot.NewTs3BotClient(grpcClient)
	if ts3Client == nil {
		return fmt.Errorf("Failed to create TS3 Bot Client")
	}

	return
}

func Disconnect() {
	if grpcClient != nil {
		grpcClient.Close()
	}
}
