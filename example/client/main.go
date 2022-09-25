package main

import (
	"log"

	"github.com/zehlt/prott"
	"github.com/zehlt/prott/example/common"
)

func main() {
	common.RegisterMessages()

	l, err := prott.Dial(prott.TCP, ":9192")
	if err != nil {
		panic(err)
	}

	log.Println("client connected")

	mess, err := l.Read()
	if err != nil {
		panic(err)
	}
	log.Printf("type: %T, value: %v\n", mess, mess)

	mess, err = l.Read()
	if err != nil {
		panic(err)
	}
	log.Printf("type: %T, value: %v\n", mess, mess)

	mess, err = l.Read()
	if err != nil {
		panic(err)
	}
	log.Printf("type: %T, value: %v\n", mess, mess)

	mess, err = l.Read()
	if err != nil {
		panic(err)
	}
	log.Printf("type: %T, value: %v\n", mess, mess)

	l.Close()

	log.Println("client disconnected")
}
