package prott

type Env struct {
	Req Request
	Res Response
}

type Request struct {
	Id     int
	Addr   string
	Packet Packet
}

type Response struct {
	id          int
	messageChan chan<- Message
}

func (r *Response) Reply(p Packet) {
	r.messageChan <- Message{T: UNICAST_MESSAGE, P: p, Receiver: r.id}
}

func (r *Response) Unicast(id int, p Packet) {
	r.messageChan <- Message{T: UNICAST_MESSAGE, P: p, Receiver: id}
}

func (r *Response) Boardcast(p Packet) {
	r.messageChan <- Message{T: BROADCAST_MESSAGE, P: p, Sender: r.id}
}

func (r *Response) Emit(p Packet) {
	r.messageChan <- Message{T: EMIT_MESSAGE, P: p}
}