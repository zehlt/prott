package main

import (
	"time"

	"github.com/zehlt/prott/client"
)

type GameClient struct {
	socket client.Socket
}

func NewGameClient() *GameClient {
	return &GameClient{
		socket: client.NewTcpSocket(),
	}
}

func main() {
	port := ":8091"

	gc := NewGameClient()
	_, err := gc.socket.Connect(port)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 5)
	// gc.socket.Disconnect()

	// fmt.Println("server started on port", port)
	// gms := NewGameServer()

	// gms.server.Register(network.USER_CONNECTED_PACKET, gms.UserConnected)
	// gms.server.Register(network.USER_DISCONNECTED_PACKET, gms.UserDisconnected)

	// waiter, err := server.NewTcpWaiter(port)
	// if err != nil {
	// 	panic(err)
	// }
	// gms.server.Serve(context.Background(), waiter)
}
