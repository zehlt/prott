package main

import (
	"fmt"

	"github.com/zehlt/prott"
)

func (g *GameServer) UserConnected(env prott.Env) {
	// env.Res.Reply(prott.Packet{T: prott.SERVER_CHAT_PACKET, Data: prott.ServerChatPacket{Message: "hello world!"}})
}

func (g *GameServer) UserDisconnected(env prott.Env) {
}

func (g *GameServer) UserChat(env prott.Env) {
	mess := env.Req.Packet().Data.(prott.UserChatPacket)
	fmt.Println("server received packer", mess.Message, env.Req.Id())

	env.Res.Unicast(env.Req.Id(), prott.Packet{T: prott.SERVER_CHAT_PACKET, Data: prott.ServerChatPacket{Message: fmt.Sprintf("%d said %s", env.Req.Id(), mess.Message)}})
	// env.Res.Emit(prott.Packet{T: prott.SERVER_CHAT_PACKET, Data: prott.ServerChatPacket{Message: fmt.Sprintf("%d said %s", env.Req.Id(), mess.Message)}})
	// env.Res.Boardcast(prott.Packet{T: prott.SERVER_CHAT_PACKET, Data: prott.ServerChatPacket{Message: fmt.Sprintf("%d said %s", env.Req.Id, mess.Message)}})
}
