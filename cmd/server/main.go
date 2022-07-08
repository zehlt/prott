package main

import (
	"context"
	"fmt"

	"github.com/zehlt/prott/network"
	"github.com/zehlt/prott/server"
)

type GameServer struct {
	router server.Router
}

func NewGameServer() *GameServer {
	return &GameServer{
		router: server.NewRouter(),
	}
}

func main() {
	network.Register()

	port := ":8091"

	fmt.Println("server started on port", port)
	gms := NewGameServer()

	gms.router.Register(network.USER_CONNECTED_PACKET, gms.UserConnected)
	gms.router.Register(network.USER_DISCONNECTED_PACKET, gms.UserDisconnected)

	waiter, err := server.NewTcpWaiter(port)
	if err != nil {
		panic(err)
	}
	gms.router.Serve(context.Background(), waiter)
}
