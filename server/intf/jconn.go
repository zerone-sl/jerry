package intf

import "net"

type XConn interface {
	Do()

	Close()

	Conn() *net.TCPConn

	RemoteAddr() string

	ConnID() string

	ConnType() string

	Send(msg []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
