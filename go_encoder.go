package prott

import (
	"bytes"
	"encoding/gob"
)

type gobEncoder struct {
	buffReader []byte
}

func newGobEncoder() Encoder {
	return &gobEncoder{
		buffReader: make([]byte, 1000),
	}
}

func (ge *gobEncoder) Register(p Packet) {
	gob.Register(p)
}

func (ge *gobEncoder) Encode(p Packet) ([]byte, error) {
	b := bytes.Buffer{}

	e := gob.NewEncoder(&b)
	err := e.Encode(p)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (ge *gobEncoder) Decode(data []byte) (Packet, error) {
	var p Packet

	b := bytes.Buffer{}
	b.Write(data)

	e := gob.NewDecoder(&b)
	err := e.Decode(&p)
	if err != nil {
		return Packet{}, err
	}

	return p, nil
}
