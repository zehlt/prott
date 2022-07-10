package main

import (
	"fmt"

	"github.com/zehlt/prott/network"
	"github.com/zehlt/prott/router"
)

func (g *GameServer) UserConnected(env router.Env) {
	env.Res.Reply(network.Packet{T: network.SERVER_CHAT_PACKET, Data: network.ServerChatPacket{Message: "hello world!"}})
}

func (g *GameServer) UserDisconnected(env router.Env) {
}

func (g *GameServer) UserChat(env router.Env) {
	mess := env.Req.Packet.Data.(network.UserChatPacket)

	// env.Res.Emit(network.Packet{T: network.SERVER_CHAT_PACKET, Data: network.ServerChatPacket{Message: fmt.Sprintf("%d said %s", env.Req.Id, mess.Message)}})
	// env.Res.Boardcast(network.Packet{T: network.SERVER_CHAT_PACKET, Data: network.ServerChatPacket{Message: fmt.Sprintf("%d said %s", env.Req.Id, mess.Message)}})
	env.Res.Unicast(env.Req.Id, network.Packet{T: network.SERVER_CHAT_PACKET, Data: network.ServerChatPacket{Message: fmt.Sprintf("%d said %s", env.Req.Id, mess.Message)}})
}
