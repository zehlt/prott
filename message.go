package prott

type PacketType int

const (
	EMIT_PACKET PacketType = iota
	UNICAST_PACKET
	MULTICAST_PACKET
	BROADCAST_PACKET
)

type Message struct {
	Data string
}

type Packet struct {
	Type     PacketType
	Sender   string
	Receiver string

	Data Message
}
