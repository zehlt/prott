package prott

type Connection interface {
	Read() (Packet, error)
	Write(p Packet) error
	LocalAddr() string
	RemoteAddr() string
	Id() int
	Close() error
}
