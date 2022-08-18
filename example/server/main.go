package main

import (
	"context"
	"log"
	"time"

	"github.com/zehlt/prott"
)

func main() {
	log.Println("hello from server")

	l, err := prott.NewTcpListener(":8090")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(l)

	r := prott.NewRouter()
	r.Serve(context.Background(), l)

	time.Sleep(5 * time.Second)

	r.Send(prott.Packet{
		Type: prott.EMIT_PACKET,
		Data: prott.Message{},
	})

	time.Sleep(1 * time.Hour)
	log.Println("SERVER CLOSED")
}
