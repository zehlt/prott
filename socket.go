package prott

type Socket interface {
	Connect(port string) (Connection, error)
	// Send(p Packet) error
	// Recv() (Packet, error)
	// Disconnect() error
}
