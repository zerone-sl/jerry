package jnet

import (
	"fmt"
	int2 "github.com/zerone/jerry/server/intf"
	"log"
	"net"
	"sync"
)

type JServerNet struct {
	Name   string
	TCPVer string
	Ip     string
	Port   int
}

func (jsn *JServerNet) Start() {
	log.Printf("[start] server listen ip=%s port=%d is starting...\n", jsn.Ip, jsn.Port)
	addr, err := net.ResolveTCPAddr(jsn.TCPVer, fmt.Sprintf("%s:%d", jsn.Ip, jsn.Port))
	if err != nil {
		log.Fatalln("resolve tcp addr err: ", err)
		return
	}

	listener, err := net.ListenTCP(jsn.TCPVer, addr)
	if err != nil {
		log.Fatalln("listen tcp addr err", err)
		return
	}

	go func() {
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Fatalln("accept tcp err: ", err)
				continue
			}

			tcpXconn := NewTCPXConn("1", conn, CallBack)
			go tcpXconn.Do()
		}
	}()
}

func (jsn *JServerNet) Serve() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()
	jsn.Start()
}

func (jsn *JServerNet) Stop() {

}

func New(name string) int2.JServer {
	return &JServerNet{Name: name, TCPVer: "tcp4", Ip: "127.0.0.1", Port: 6069}
}
