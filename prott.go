package prott

func NewTcpSocket() Socket {
	return newTcpSocket()
}

func NewTcpWaiter(port string) (Waiter, error) {
	return newTcpWaiter(port)
}

func NewRouter() Router {
	return newRouter()
}
