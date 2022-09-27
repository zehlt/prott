package prott

import (
	"net"

	"github.com/zehlt/datt"
)

type Conn interface {
	Read() (Message, error)
	Write(m Message) error
	Close() error
	// deadlines
	// Addr()
}

type tcpConn struct {
	conn net.Conn

	wbuff [1600]byte

	rbuff     [1600]byte
	readQueue *datt.Queue[Message]
}

func newTcpConn(conn net.Conn) Conn {
	return &tcpConn{
		conn:      conn,
		readQueue: datt.NewQueue[Message](),
	}
}

func (c *tcpConn) Read() (Message, error) {
	if !c.readQueue.IsEmpty() {
		message, _ := c.readQueue.Dequeue()
		return message, nil
	}

	// n, err := c.conn.Read(c.rbuff[:])
	// if err != nil {
	// 	return nil, err
	// }

	// rs := ReadStream{buff: c.rbuff[:n]}

	// for n > rs.TotalRead() {
	// 	// var t uint16

	// 	// err := rs.ReadUint16(&t)()
	// 	// if err != nil {
	// 	// 	return nil, err
	// 	// }

	// 	// message, err := getMessageOfType(MessageType(t))
	// 	// if err != nil {
	// 	// 	return nil, err
	// 	// }

	// 	// err = message.Read(&rs)
	// 	// if err != nil {
	// 	// 	return nil, err
	// 	// }

	// 	// c.readQueue.Enqueue(message)
	// }

	m, _ := c.readQueue.Dequeue()
	return m, nil
}

func (c *tcpConn) Write(m Message) error {
	// t := m.Type()

	// ws := WriteStream{buff: c.wbuff[:]}

	// err := ws.WriteUint16(uint16(t))()
	// if err != nil {
	// 	return err
	// }

	// err = m.Write(&ws)
	// if err != nil {
	// 	return err
	// }

	// n, err := c.conn.Write(c.wbuff[:ws.TotalWrite()])
	// if err != nil {
	// 	return err
	// }

	// // TODO: check if necessary
	// if n != ws.total {
	// 	panic("TODO: N and WS.TOTAL isn't equal")
	// }

	return nil
}

func (c *tcpConn) Close() error {
	return c.conn.Close()
}
