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

type SocketMessageType int

const (
	SEND_MESSAGE_SOCKET SocketMessageType = iota
	CLOSE_MESSAGE_SOCKET
)

type SocketMessage struct {
	T SocketMessageType
	P Packet
}
