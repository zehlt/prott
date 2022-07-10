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
	fmt.Println("connection completly")

	fmt.Println("TRY TO SEND")
	bus.Send(network.Packet{T: network.USER_CHAT_PACKET, Data: network.UserChatPacket{Message: "hello from me"}})
	time.Sleep(time.Second * 1)

	for i := 0; i < 100; i++ {
		data, ok := bus.TryRecv()
		if ok {
			fmt.Println("message", data)
		}
	}

	// bus.Send(network.Packet{T: network.USER_CHAT_PACKET, Data: network.UserChatPacket{Message: "hello from me"}})
	time.Sleep(time.Millisecond * 2)

	gc.socket.Disconnect()
	fmt.Println("disconnected completly")

	// for i := 0; i < 10; i++ {
	// 	bus, err := gc.socket.Connect(port)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("connection completly")

	// 	data := bus.Recv()
	// 	fmt.Println("message", data)
	// 	time.Sleep(time.Millisecond * 500)

	// 	bus.Send(network.Packet{T: network.USER_CHAT_PACKET, Data: network.UserChatPacket{Message: "hello from me"}})
	// 	time.Sleep(time.Millisecond * 30)

	// 	bus.Send(network.Packet{T: network.USER_CHAT_PACKET, Data: network.UserChatPacket{Message: "hello from me"}})
	// 	time.Sleep(time.Millisecond * 90)

	// 	bus.Send(network.Packet{T: network.USER_CHAT_PACKET, Data: network.UserChatPacket{Message: "hello from me"}})
	// 	time.Sleep(time.Millisecond * 10)

	// 	bus.Send(network.Packet{T: network.USER_CHAT_PACKET, Data: network.UserChatPacket{Message: "hello from me"}})
	// 	time.Sleep(time.Millisecond * 500)

	// 	gc.socket.Disconnect()
	// 	fmt.Println("disconnected completly")

	// 	time.Sleep(time.Millisecond * 500)
	// }
}
