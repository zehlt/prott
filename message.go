package prott

type MessageType int

type Message interface {
	Type() MessageType

	// Write(ws *WriteStream) error
	// Read(rs *ReadStream) error
}

// type MessageBufferOperation func() error

// func Operations(ops ...MessageBufferOperation) error {

// 	for _, operation := range ops {
// 		err := operation()
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// type WriteStream struct {
// 	buff   []byte
// 	cursor int
// 	total  int
// }

// func (ws *WriteStream) WriteGob(v any) func() error {
// 	return func() error {

// 		// buf := bytes.Buffer{}
// 		// e := gob.NewEncoder(&buf)

// 		// buf.Len()
// 		// e.Encode()

// 		// size, err := checkSize(ws.buff, ws.cursor, 2)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// binary.LittleEndian.PutUint16(ws.buff[ws.cursor:], v)
// 		// ws.total += size
// 		// ws.cursor += size
// 		return nil
// 	}
// }

// func (ws *WriteStream) WriteUint16(v uint16) func() error {
// 	return func() error {
// 		size, err := checkSize(ws.buff, ws.cursor, 2)
// 		if err != nil {
// 			return err
// 		}
// 		binary.LittleEndian.PutUint16(ws.buff[ws.cursor:], v)
// 		ws.total += size
// 		ws.cursor += size
// 		return nil
// 	}
// }

// func (ws *WriteStream) TotalWrite() int {
// 	return ws.total
// }

// type ReadStream struct {
// 	buff   []byte
// 	cursor int
// 	total  int
// }

// func (rs *ReadStream) ReadUint16(v *uint16) func() error {
// 	return func() error {
// 		size, err := checkSize(rs.buff, rs.cursor, 2)
// 		if err != nil {
// 			return err
// 		}
// 		*v = binary.LittleEndian.Uint16(rs.buff[rs.cursor:])
// 		rs.total += size
// 		rs.cursor += size
// 		return nil
// 	}
// }

// func (rs *ReadStream) TotalRead() int {
// 	return rs.total
// }

// func checkSize(b []byte, cursor int, expected int) (int, error) {
// 	size := len(b) - cursor

// 	if size < expected {
// 		return 0, fmt.Errorf("wrong message packet")
// 	}

// 	return expected, nil
// }
