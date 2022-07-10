package prott

type statusType int

const (
	CONNECTION_STATUS statusType = iota
	DISCONNECTION_STATUS
)

type status struct {
	t          statusType
	connection Connection
}
