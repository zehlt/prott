package prott

type messageType int

const (
	EMIT_MESSAGE messageType = iota
	UNICAST_MESSAGE
	MULTICAST_MESSAGE
	BROADCAST_MESSAGE
)

type message struct {
	t        messageType
	p        Packet
	sender   int
	receiver int
}
