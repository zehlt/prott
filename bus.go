package prott

type Bus struct {
	recvC <-chan Packet
	sendC chan<- Packet
}

func (b *Bus) Send(p Packet) {
	b.sendC <- p
}

func (b *Bus) TryRecv() (Packet, bool) {
	if len(b.recvC) <= 0 {
		return Packet{}, false
	}
	return <-b.recvC, true
}

func (b *Bus) Recv() Packet {
	return <-b.recvC
}
