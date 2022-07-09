package server

import "github.com/zehlt/prott/network"

type StatusType int

const (
	CONNECTION_STATUS StatusType = iota
	DISCONNECTION_STATUS
)

type Status struct {
	T          StatusType
	Connection network.Connection
}
