package main

import (
	"fmt"
	"time"

	"github.com/zehlt/prott/network"
	"github.com/zehlt/prott/socket"
)

type GameClient struct {
	socket socket.Socket
}

func NewGameClient() *GameClient {
	return &GameClient{
		socket: socket.NewTcpSocket(),
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
	time.Sleep(time.Second * 2)

	fmt.Println("TRY TO SEND")
	bus.Send(network.Packet{T: network.USER_CHAT_PACKET, Data: network.UserChatPacket{Message: "hello from me"}})

	for i := 0; i < 2; i++ {
		data := bus.Recv()
		fmt.Println("message", data)
	}

	// bus.Send(network.Packet{T: network.USER_CHAT_PACKET, Data: network.UserChatPacket{Message: "hello from me"}})
	time.Sleep(time.Second * 3)

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
