package prott

type StatusType int

const (
	CONNECTION_STATUS StatusType = iota
	DISCONNECTION_STATUS
)

type Status struct {
	T          StatusType
	Connection Connection
}
