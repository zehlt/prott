package main

import (
	"fmt"
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
	bus, err := gc.socket.Connect(port)
	if err != nil {
		panic(err)
	}

	data := bus.Recv()
	fmt.Println(data)

	time.Sleep(time.Second * 3)

	gc.socket.Disconnect()
	fmt.Println("disconnected completly")

	time.Sleep(time.Second * 3)

	fmt.Println("retry connection")

	bus, err = gc.socket.Connect(port)
	if err != nil {
		panic(err)
	}

	data = bus.Recv()
	fmt.Println(data)

	time.Sleep(time.Second * 3)

}
