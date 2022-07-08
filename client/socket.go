package client

type Socket interface {
	Connect(port string) (Bus, error)
	Disconnect()
}
