package prott

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type Router struct {
	r routeTable[string]
}

func NewRouter() *Router {
	var s Router
	s.r = newRouteTable[string]()
	return &s
}

func (s *Router) Send(p Packet) {

	switch p.Type {
	case EMIT_PACKET:
		routes := s.r.GetAll()

		for _, route := range routes {
			log.Println(route)
			route.c <- p
		}

		log.Println("EMIT")

	case MULTICAST_PACKET:
		log.Println("MULT")

	case BROADCAST_PACKET:
		log.Println("BROA")

	case UNICAST_PACKET:
		log.Println("UNIC")
	}
}

func (s *Router) Serve(ctx context.Context, l Listener) {
	handleClientConnection(ctx, registerClientConnection(ctx, &s.r, waitForConnection(ctx, l)))
}

func waitForConnection(ctx context.Context, l Listener) <-chan Connection {
	out := make(chan Connection)

	go func() {
		defer close(out)

		for {
			conn, err := l.Accept()
			if err != nil {
				log.Fatal("error in wait for connection")
			}

			out <- conn
		}
	}()

	return out
}

type RegisteredConnection struct {
	Send chan Packet
	Conn Connection
}

func registerClientConnection(ctx context.Context, routeTable *routeTable[string], connections <-chan Connection) <-chan RegisteredConnection {
	out := make(chan RegisteredConnection)

	go func() {
		defer close(out)

		for connection := range connections {
			log.Println("CONNECTION: ", connection.RemoteAddr(), connection.LocalAddr())

			send := make(chan Packet)
			routeTable.Add(connection.RemoteAddr(), Route{c: send, addr: connection.RemoteAddr()})

			out <- RegisteredConnection{
				Send: send,
				Conn: connection,
			}
		}

	}()

	return out
}

// func middlewarePacket(ctx context.Context, packets <-chan Packet, fn func(Packet) (Packet, bool)) <-chan Packet {
// 	out := make(chan Packet)

// 	go func() {
// 		defer close(out)

// 		for packet := range packets {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			default:
// 				p, ok := fn(packet)
// 				if ok {
// 					out <- p
// 				}
// 			}
// 		}
// 	}()

// 	return out
// }

func handleClientConnection(ctx context.Context, connections <-chan RegisteredConnection) <-chan Packet {
	packets := make(chan Packet)
	// defer close(packets)

	go func() {
		for connection := range connections {

			// read
			go func(conn RegisteredConnection) {
				for {
					p, err := conn.Conn.Read()
					if err != nil {
						// TODO: do some error handling
						// connection closed
						log.Fatal(err)
					}

					packets <- p
				}
			}(connection)

			// send
			go func(conn RegisteredConnection) {
				for p := range conn.Send {
					log.Println("GONNA SEND:", p)

					err := conn.Conn.Write(p)
					if err != nil {
						// TODO: do some error handling

						log.Fatal(err)
					}
				}
			}(connection)
		}
	}()

	return packets
}

// ROUTE TABLE

type Route struct {
	c    chan<- Packet
	addr string
}

type routeTable[T comparable] struct {
	routes map[T]Route

	rwMutex sync.RWMutex
}

func newRouteTable[T comparable]() routeTable[T] {
	return routeTable[T]{
		routes: make(map[T]Route),
	}
}

func (r *routeTable[T]) Add(id T, c Route) {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	log.Println("ADDIND A ROUTE")

	r.routes[id] = c
}

func (r *routeTable[T]) Remove(id T) {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	delete(r.routes, id)
}

func (r *routeTable[T]) Get(id T) (Route, error) {
	r.rwMutex.RLock()
	defer r.rwMutex.RUnlock()

	c, ok := r.routes[id]
	if !ok {
		return Route{}, fmt.Errorf("id does not match")
	}

	return c, nil
}

func (r *routeTable[T]) GetAll() []Route {
	r.rwMutex.RLock()
	defer r.rwMutex.RUnlock()

	arr := make([]Route, len(r.routes))

	i := 0
	for _, v := range r.routes {
		arr[i] = v
		i++
	}

	return arr
}
