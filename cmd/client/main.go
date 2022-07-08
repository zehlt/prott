package main

import (
	"time"

	"github.com/zehlt/prott/client"
	"github.com/zehlt/prott/network"
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
	network.Register()

	port := ":8091"

	gc := NewGameClient()
	_, err := gc.socket.Connect(port)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 5)

	gc.socket.Disconnect()

	time.Sleep(time.Second * 3)
}
