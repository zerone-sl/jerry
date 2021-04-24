package intf

type XReq interface {
	Conn() XConn

	Msg() []byte
}
