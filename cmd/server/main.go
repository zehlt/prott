package main

import (
	"context"
	"fmt"

	"github.com/zehlt/prott"
)

type GameServer struct {
	router prott.Router
}

func NewGameServer() *GameServer {
	return &GameServer{
		router: prott.NewRouter(),
	}
}

func main() {
	prott.Register()

	port := ":8091"

	fmt.Println("router started on port", port)
	gms := NewGameServer()

	gms.router.Register(prott.USER_CONNECTED_PACKET, gms.UserConnected)
	gms.router.Register(prott.USER_DISCONNECTED_PACKET, gms.UserDisconnected)
	gms.router.Register(prott.USER_CHAT_PACKET, gms.UserChat)

	waiter, err := prott.NewTcpWaiter(port)
	if err != nil {
		panic(err)
	}
	gms.router.Serve(context.Background(), waiter)
}
