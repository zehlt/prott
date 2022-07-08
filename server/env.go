package server

import "github.com/zehlt/prott/network"

type Env struct {
	Req Request
	Res Response
}

type Request struct {
	Id     int
	Addr   string
	Packet network.Packet
}

type Response struct {
}

func (r *Response) Unicast() {

}

func (r *Response) Boardcast() {

}

func (r *Response) Emit() {

}
