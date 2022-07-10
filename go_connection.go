package prott

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

type goConnection struct {
	conn       net.Conn
	buffReader []byte
	id         int
}

func newGoConnection(conn net.Conn, id int) Connection {

	return &goConnection{
		conn:       conn,
		buffReader: make([]byte, 1000),
		id:         id,
	}
}

func (c *goConnection) Read() (Packet, error) {
	// c.conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	buffer := bytes.Buffer{}

	// Check if the best way to do it
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

	p, err := gobDecodePacket(buffer)
	if err != nil {
		return Packet{}, err
	}

	return p, nil
}

func (c *goConnection) Write(p Packet) error {

	b, err := gobEncodePacket(p)
	if err != nil {
		return err
	}

	// TODO: check if size is important
	_, err = c.conn.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (c *goConnection) LocalAddr() string {
	return c.conn.LocalAddr().String()
}

func (c *goConnection) RemoteAddr() string {
	return c.conn.RemoteAddr().String()
}

func (c *goConnection) Id() int {
	return c.id
}

func (c *goConnection) Close() error {
	err := c.conn.Close()
	if err != nil {
		return fmt.Errorf("unable to close the server correctly: %s", err)
	}

	return nil
}

func gobDecodePacket(buffer bytes.Buffer) (Packet, error) {
	var p Packet

	e := gob.NewDecoder(&buffer)
	err := e.Decode(&p)
	if err != nil {
		return Packet{}, err
	}

	return p, nil
}

func gobEncodePacket(p Packet) ([]byte, error) {
	b := bytes.Buffer{}

	e := gob.NewEncoder(&b)
	err := e.Encode(p)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
