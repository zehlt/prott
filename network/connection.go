package network

type Connection interface {
	Read() (Packet, error)
	Write(p Packet) error
	Addr() string
	Id() int
	Close() error
}
