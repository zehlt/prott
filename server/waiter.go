package server

import "github.com/zehlt/prott/network"

type Waiter interface {
	Accept() (network.Connection, error)
	Close() error
	Addr() string
}
