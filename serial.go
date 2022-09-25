package prott

import (
	"fmt"
	"reflect"
	"sync"
)

var serial = map[MessageType]reflect.Type{}
var mu = sync.Mutex{}

func RegisterMessage(m Message) {
	mu.Lock()
	defer mu.Unlock()

	if m == nil {
		panic("message cannot be nil")
	}

	_, ok := serial[m.Type()]
	if ok {
		panic("message type already registered")
	}

	t := reflect.ValueOf(m).Elem().Type()
	serial[m.Type()] = t
}

func getMessageOfType(t MessageType) (Message, error) {
	m, ok := serial[t]
	if !ok {
		return nil, fmt.Errorf("message of type t does not exist")
	}

	newMessage := reflect.New(m).Interface().(Message)
	return newMessage, nil
}

// func hasMessageOfType(t MessageType) bool {
// 	_, ok := serial[t]
// 	return ok
// }
