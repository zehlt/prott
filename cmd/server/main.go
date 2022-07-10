package main

import (
	"context"
	"fmt"

	"github.com/zehlt/prott/network"
	"github.com/zehlt/prott/router"
)

type GameServer struct {
	router router.Router
}

func NewGameServer() *GameServer {
	return &GameServer{
		router: router.NewRouter(),
	}
}

func main() {
	network.Register()

	port := ":8091"

	fmt.Println("router started on port", port)
	gms := NewGameServer()

	gms.router.Register(network.USER_CONNECTED_PACKET, gms.UserConnected)
	gms.router.Register(network.USER_DISCONNECTED_PACKET, gms.UserDisconnected)
	gms.router.Register(network.USER_CHAT_PACKET, gms.UserChat)

	waiter, err := router.NewTcpWaiter(port)
	if err != nil {
		panic(err)
	}
	gms.router.Serve(context.Background(), waiter)
}
