package jnet

import "github.com/zerone/jerry/server/intf"

type TCPXReq struct {
	conn *TCPXConn

	msg []byte
}

func (tr *TCPXReq) Conn() intf.XConn {
	return tr.conn
}

func (tr *TCPXReq) Msg() []byte {
	return tr.msg
}

func NewTCPXReq(conn *TCPXConn, msg []byte) intf.XReq {
	tcpXreq := &TCPXReq{
		conn: conn,
		msg:  msg,
	}
	return tcpXreq
}
