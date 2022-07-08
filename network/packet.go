package network

import "encoding/gob"

type PacketType int

const (
	USER_CONNECTED_PACKET PacketType = iota
	USER_DISCONNECTED_PACKET

	SERVER_CHAT_PACKET

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

type ServerChatPacket struct {
	Message string
}

func Register() {
	gob.Register(Packet{})
	gob.Register(UserConnectedPacket{})
	gob.Register(UserDisconnectedPacket{})
	gob.Register(ServerChatPacket{})
	gob.Register(ServerConnectionAcceptedPacket{})
}
