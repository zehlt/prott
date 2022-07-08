package network

type PacketType int

const (
	USER_CONNECTED_PACKET PacketType = iota
	USER_DISCONNECTED_PACKET

	SERVER_CONNECTION_ACCEPTED_PACKET
	CUSTOM_PACKET
)

type Packet struct {
	T    PacketType
	Data interface{}
}

type UserDisconnectedPacket struct {
}

type UserConnectedPacket struct {
}

type ServerConnectionAcceptedPacket struct {
}
