package prott

import (
	"bytes"
	"fmt"
	"net"
)

type goConnection struct {
	conn net.Conn
	enc  Encoder

	buffReader []byte
}

func newGoConnection(conn net.Conn) Connection {

	return &goConnection{
		conn: conn,
		enc:  newGobEncoder(),

		buffReader: make([]byte, 1000),
	}
}

func (c *goConnection) Read() (Packet, error) {
	buffer := bytes.Buffer{}

	// TODO: check if deadline is neccesary
	// c.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	before, err := c.conn.Read(c.buffReader)
	if err != nil {
		return Packet{}, err
	}

	after, err := buffer.Write(c.buffReader[:before])
	if err != nil {
		return Packet{}, err
	}

	if before != after {
		return Packet{}, fmt.Errorf("unable to write inside the new buffer")
	}

	p, err := c.enc.Decode(buffer.Bytes())
	if err != nil {
		return Packet{}, err
	}

	return p, nil
}

func (c *goConnection) Write(p Packet) error {

	b, err := c.enc.Encode(p)
	if err != nil {
		return err
	}

	// TODO: check if size is important
	_, err = c.conn.Write(b)
	return err
}

func (c *goConnection) LocalAddr() string {
	return c.conn.LocalAddr().String()
}

func (c *goConnection) RemoteAddr() string {
	return c.conn.RemoteAddr().String()
}

func (c *goConnection) Close() error {
	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("unable to close the server correctly: %s", err)
	}

	return nil
}
