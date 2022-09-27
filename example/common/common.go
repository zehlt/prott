package common

import "github.com/zehlt/prott"

const (
	TRANSFORM_MESSAGE prott.MessageType = iota
	MOVE_REQUEST_MESSAGE
)

func RegisterMessages() {
	prott.RegisterMessage(&TransformMessage{})
	prott.RegisterMessage(&MoveRequestMessage{})
}

type TransformMessage struct {
	X uint16
	Y uint16
	Z uint16
}

func (p *TransformMessage) Type() prott.MessageType {
	return TRANSFORM_MESSAGE
}

// func (p *TransformMessage) Write(ws *prott.WriteStream) error {
// 	return prott.Operations(
// 		ws.WriteUint16(p.X),
// 		ws.WriteUint16(p.Y),
// 		ws.WriteUint16(p.Z),
// 	)
// }

// func (p *TransformMessage) Read(rs *prott.ReadStream) error {
// 	return prott.Operations(
// 		rs.ReadUint16(&p.X),
// 		rs.ReadUint16(&p.Y),
// 		rs.ReadUint16(&p.Z),
// 	)
// }

type MoveRequestMessage struct {
	X uint16
	Y uint16
}

func (p *MoveRequestMessage) Type() prott.MessageType {
	return MOVE_REQUEST_MESSAGE
}

// func (p *MoveRequestMessage) Write(ws *prott.WriteStream) error {
// 	return prott.Operations(
// 		ws.WriteUint16(p.X),
// 		ws.WriteUint16(p.Y),
// 	)
// }

// func (p *MoveRequestMessage) Read(rs *prott.ReadStream) error {
// 	return prott.Operations(
// 		rs.ReadUint16(&p.X),
// 		rs.ReadUint16(&p.Y),
// 	)
// }
