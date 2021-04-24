package jnet

import (
	"errors"
	"github.com/zerone/jerry/server/intf"
	"log"
	"net"
)

type TCPXConn struct {
	connType string

	connID string

	conn *net.TCPConn

	isClosed bool

	ExitChan chan bool

	Router intf.XRouter
}

func NewTCPXConn(connID string, conn *net.TCPConn, router intf.XRouter) *TCPXConn {
	tpcXconn := &TCPXConn{
		connType: "tcp4",
		connID:   connID,
		conn:     conn,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
	return tpcXconn
}

func (t *TCPXConn) Do() {
	Read(t)
}

func (t *TCPXConn) Close() {
	if !t.isClosed {
		t.conn.Close()
		close(t.ExitChan)
		t.isClosed = true
	}
}

func (t *TCPXConn) Conn() *net.TCPConn {
	return t.conn
}

func (t *TCPXConn) RemoteAddr() string {
	return ""
}

func (t *TCPXConn) ConnID() string {
	return t.connID
}

func (t *TCPXConn) ConnType() string {
	return t.connType
}

func (t *TCPXConn) Send(msg []byte) error {
	return nil
}

func CallBack(conn *net.TCPConn, msg []byte, cnt int) error {
	log.Println("callback client ...")
	if _, err := conn.Write(msg[:cnt]); err != nil {
		log.Println("callback to client err: ", err)
		return errors.New("callback to client err")
	}
	return nil
}

func Read(t *TCPXConn) {
	for {
		buf := make([]byte, 512)
		_, err := t.conn.Read(buf)
		if err != nil {
			continue
		}

		xreq := NewTCPXReq(t, buf)

		go func(r intf.XReq) {
			t.Router.PreHandle(r)
			t.Router.Handle(r)
			t.Router.PostHandle(r)
		}(xreq)
	}
}
