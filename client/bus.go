package client

import "github.com/zehlt/prott/network"

type Bus struct {
	recv <-chan network.Packet
	send chan<- network.Packet
}

func (b *Bus) Send(p network.Packet) {
	b.send <- p
}

func (b *Bus) Recv() (network.Packet, bool) {
	if len(b.recv) <= 0 {
		return network.Packet{}, false
	}
	return <-b.recv, true
}
