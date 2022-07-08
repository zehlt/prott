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

func (g *GameServer) UserConnected(env server.Env) {
	fmt.Println("user connected", env.Req.Id, env.Req.Addr)
}

func (g *GameServer) UserDisconnected(env server.Env) {
	fmt.Println("user disconnected", env.Req.Id, env.Req.Addr)
}
