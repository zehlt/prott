package main

import (
	"log"

	"github.com/zehlt/prott"
	"github.com/zehlt/prott/example/common"
)

func main() {
	common.RegisterMessages()

	l, err := prott.Listen(prott.TCP, ":9192")
	if err != nil {
		panic(err)
	}

	log.Println("starting server")

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		log.Println("client connected to server")
		go handleConnection(conn)
	}
}

func handleConnection(c prott.Conn) {
	defer c.Close()

	err := c.Write(&common.TransformMessage{X: 111, Y: 2, Z: 984})
	if err != nil {
		panic(err)
	}

	// time.Sleep(50 * time.Microsecond)

	err = c.Write(&common.MoveRequestMessage{X: 222, Y: 698})
	if err != nil {
		panic(err)
	}

	err = c.Write(&common.MoveRequestMessage{X: 333, Y: 698})
	if err != nil {
		panic(err)
	}

	err = c.Write(&common.MoveRequestMessage{X: 444, Y: 698})
	if err != nil {
		panic(err)
	}

	for {
		message, err := c.Read()
		if err != nil {
			log.Println("client disconnected to server")
			return
		}

		log.Println(message)
	}
}
