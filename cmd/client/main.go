package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/zehlt/prott"
)

type GameClient struct {
	socket prott.Socket
}

func NewGameClient() *GameClient {
	return &GameClient{
		socket: prott.NewTcpSocket(),
	}
}

func main() {
	prott.Register()

	port := ":8091"
	gc := NewGameClient()

	bus, err := gc.socket.Connect(port)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	bus.Register(prott.SERVER_CLOSE_CONNECTION_PACKET, func(p prott.Packet) {
		bus.Close()
		cancel()
	})

	bus.Register(prott.SERVER_CHAT_PACKET, func(p prott.Packet) {
		fmt.Println("MESSAGE RECEIVED")
	})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

	send:
		for {
			select {
			case <-ctx.Done():
				break send
			default:
				bus.Send(prott.Packet{T: prott.USER_CHAT_PACKET, Data: prott.UserChatPacket{Message: "hello from me"}})
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	go func() {
		defer wg.Done()

	read:
		for {
			select {
			case <-ctx.Done():
				break read
			default:
				bus.Recv()
			}
		}
	}()

	go func() {
		time.Sleep(time.Millisecond * 6500)
		bus.Close()
		cancel()
	}()

	wg.Wait()
	fmt.Println("CLIENT CLOSED")
}
