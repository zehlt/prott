package prott

type Waiter interface {
	Accept() (Connection, error)
	Close() error
	Addr() string
}
