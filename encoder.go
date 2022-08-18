package prott

type Encoder interface {
	Register(p Packet)
	Encode(p Packet) ([]byte, error)
	Decode(data []byte) (Packet, error)
}
