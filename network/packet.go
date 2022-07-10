package network

import "encoding/gob"

type PacketType int

const (
	USER_CONNECTED_PACKET PacketType = iota
	USER_DISCONNECTED_PACKET
	USER_CHAT_PACKET

	SERVER_CHAT_PACKET

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

type UserChatPacket struct {
	Message string
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
	gob.Register(UserChatPacket{})
	gob.Register(ServerChatPacket{})
	gob.Register(ServerConnectionAcceptedPacket{})
}
