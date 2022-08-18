package main

import (
	"log"
	"time"

	"github.com/zehlt/prott"
)

func main() {
	log.Println("hello from client")

	socket := prott.NewTcpSocket()
	conn, err := socket.Connect(":8090")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() {
		p, err := conn.Read()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("RECEIVED PACKET: ", p)
	}()

	time.Sleep(30 * time.Second)
}
