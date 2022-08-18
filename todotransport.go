package prott

// type tcpSocket struct {
// 	// b    datt.DoubleBuffer[Packet]
// 	// send chan<- Packet
// 	// conn connection

// 	// isConnected bool
// }

// func (s *Transport) Send(p Packet) error {
// 	return nil
// }

// func (s *Transport) Recv() (Packet, error) {
// 	return Packet{}, nil
// }

// func (s *Transport) Disconnect() error {
// 	return nil
// }

// // func newTcpSocket() Socket {

// // 	// return &tcpSocket{
// // 	// 	b:           datt.NewDoubleBuffer[Packet](),
// // 	// 	isConnected: false,
// // 	// }
// // }

// func (c *tcpSocket) Connect(port string) error {
// 	conn, err := net.Dial("tcp", port)
// 	if err != nil {
// 		return err
// 	}

// 	c.conn = newGoConnection(conn, 0)
// 	send := c.handleServerConnection(c.conn)
// 	c.send = send

// 	c.isConnected = true
// 	return nil
// }

// func (c *tcpSocket) handleServerConnection(conn Connection) chan<- Packet {
// 	send := make(chan Packet, 1)
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go func() {
// 		defer conn.Close()

// 	loop:
// 		for {
// 			select {

// 			case <-ctx.Done():
// 				break loop

// 			case packet := <-send:
// 				err := conn.Write(packet)
// 				if err != nil {
// 					panic(fmt.Sprintln("error trying to send packet", err))
// 				}
// 			}
// 		}
// 	}()

// 	go func() {
// 		for {
// 			packet, err := conn.Read()
// 			if err != nil {
// 				c.b.PushEvent(Packet{T: SERVER_CLOSE_CONNECTION_PACKET, Data: ServerCloseConnectionPacket{Err: err.Error()}})
// 				cancel()
// 				break
// 			}
// 			c.b.PushEvent(packet)
// 		}
// 	}()

// 	return send
// }

// func (s *tcpSocket) Send(p Packet) error {
// 	if s.isConnected {
// 		s.send <- p
// 		return nil
// 	} else {
// 		return fmt.Errorf("trying to send packet on closed connection")
// 	}
// }

// func (s *tcpSocket) Recv() []Packet {
// 	s.b.SwitchBuff()
// 	var arr []Packet

// 	for {
// 		p, ok := s.b.PopEvent()
// 		if !ok {
// 			break
// 		}

// 		arr = append(arr, p)
// 	}

// 	return arr
// }

// func (s *tcpSocket) Disconnect() error {
// 	s.conn.Close()
// 	s.isConnected = false
// 	return nil
// }
