package main

import (
	"fmt"

	"github.com/zehlt/prott/network"
	"github.com/zehlt/prott/server"
)

func (g *GameServer) UserConnected(env server.Env) {
	fmt.Println("user connected", env.Req.Id, env.Req.Addr)

	env.Res.Reply(network.Packet{T: network.SERVER_CHAT_PACKET, Data: network.ServerChatPacket{Message: "hello world!"}})
}

func (g *GameServer) UserDisconnected(env server.Env) {
	fmt.Println("user disconnected", env.Req.Id, env.Req.Addr)
}

func (g *GameServer) UserChat(env server.Env) {
	mess := env.Req.Packet.Data.(network.UserChatPacket)

	env.Res.Boardcast(network.Packet{T: network.SERVER_CHAT_PACKET, Data: network.ServerChatPacket{Message: fmt.Sprintf("%d said %s", env.Req.Id, mess.Message)}})
	fmt.Println("BROADCAST DONE ")
}
