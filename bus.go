package prott

type Bus struct {
	RecvC <-chan Packet
	SendC chan<- Packet
}

func (b *Bus) Send(p Packet) {
	b.SendC <- p
}

func (b *Bus) TryRecv() (Packet, bool) {
	if len(b.RecvC) <= 0 {
		return Packet{}, false
	}
	return <-b.RecvC, true
}

func (b *Bus) Recv() Packet {
	return <-b.RecvC
}
