package prott

type Env struct {
	Req Request
	Res Response
}

type Request struct {
	id     int
	addr   string
	packet Packet
}

func (r *Request) Packet() Packet {
	return r.packet
}

func (r *Request) Id() int {
	return r.id
}

func (r *Request) Addr() string {
	return r.addr
}

type Response struct {
	id          int
	messageChan chan<- message
}

func (r *Response) Reply(p Packet) {
	r.messageChan <- message{t: UNICAST_MESSAGE, p: p, receiver: r.id}
}

func (r *Response) Unicast(id int, p Packet) {
	r.messageChan <- message{t: UNICAST_MESSAGE, p: p, receiver: id}
}

func (r *Response) Boardcast(p Packet) {
	r.messageChan <- message{t: BROADCAST_MESSAGE, p: p, sender: r.id}
}

func (r *Response) Emit(p Packet) {
	r.messageChan <- message{t: EMIT_MESSAGE, p: p}
}
