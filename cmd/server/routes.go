package main

import (
	"fmt"

	"github.com/zehlt/prott/network"
	"github.com/zehlt/prott/server"
)

func (g *GameServer) UserConnected(env server.Env) {
	env.Res.Reply(network.Packet{T: network.SERVER_CHAT_PACKET, Data: network.ServerChatPacket{Message: "hello world!"}})
}

func (g *GameServer) UserDisconnected(env server.Env) {
}

func (g *GameServer) UserChat(env server.Env) {
	mess := env.Req.Packet.Data.(network.UserChatPacket)

	// env.Res.Emit(network.Packet{T: network.SERVER_CHAT_PACKET, Data: network.ServerChatPacket{Message: fmt.Sprintf("%d said %s", env.Req.Id, mess.Message)}})
	// env.Res.Boardcast(network.Packet{T: network.SERVER_CHAT_PACKET, Data: network.ServerChatPacket{Message: fmt.Sprintf("%d said %s", env.Req.Id, mess.Message)}})
	env.Res.Unicast(env.Req.Id, network.Packet{T: network.SERVER_CHAT_PACKET, Data: network.ServerChatPacket{Message: fmt.Sprintf("%d said %s", env.Req.Id, mess.Message)}})
}
