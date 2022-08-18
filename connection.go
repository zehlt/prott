package prott

type Listener interface {
	Accept() (Connection, error)
	Close() error
	// Addr() string
}

type Connection interface {
	Read() (Packet, error)
	Write(Packet) error
	LocalAddr() string
	RemoteAddr() string
	Close() error
}
