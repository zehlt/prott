package server

// func NewUdpWaiter(port string) (Waiter, error) {
// 	l, err := net.Listen("udp", port)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &udpWaiter{l: l}, nil
// }

// type udpWaiter struct {
// 	// l net.Listener
// }

// func (w *udpWaiter) Accept() (Connection, error) {
// 	// conn, err := w.l.Accept()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// return NewGoConnection(conn), nil
// 	return nil, nil
// }

// func (w *udpWaiter) Close() error {
// 	// return w.l.Close()
// 	return nil
// }

// func (w *udpWaiter) Addr() string {
// 	// return w.l.Addr().String()
// 	return nil
// }

// UDP
// type goUdpServer struct {
// 	udpConn *net.UDPConn
// 	udpAddr *net.UDPAddr
// }

// // func newGoUdpServer() Server {
// // 	return &goUdpServer{}
// // }

// func (s *goUdpServer) Listen(port string) {
// 	addr, err := net.ResolveUDPAddr("udp", port)
// 	if err != nil {
// 		panic(err)
// 	}
// 	s.udpAddr = addr

// 	conn, err := net.ListenUDP("udp", addr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	s.udpConn = conn

// 	// buffer := make([]byte, 1024)
// 	// n, addr, err := conn.ReadFromUDP(buffer)
// 	// log.Println("-> ", string(buffer[0:n-1]), " : ", addr)
// }

// func (s *goUdpServer) Close() {
// 	s.udpConn.Close()
// }

// func (s *goUdpServer) IsOpen() bool {
// 	return false
// }
