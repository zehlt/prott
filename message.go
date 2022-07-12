package prott

type MessageType int

const (
	EMIT_MESSAGE MessageType = iota
	UNICAST_MESSAGE
	MULTICAST_MESSAGE
	BROADCAST_MESSAGE
)

type Message struct {
	T        MessageType
	P        Packet
	Sender   int
	Receiver int
}
