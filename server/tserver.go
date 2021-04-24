package jnet

import (
	"fmt"
	intf "github.com/zerone/jerry/server/intf"
	"log"
	"net"
	"sync"
)

type Server struct {
	Name   string
	TCPVer string
	Ip     string
	Port   int
	Router *intf.XRouter
}

func (s *Server) Start() {
	log.Printf("[start] server listen ip=%s port=%d is starting...\n", s.Ip, s.Port)
	addr, err := net.ResolveTCPAddr(s.TCPVer, fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		log.Fatalln("resolve tcp addr err: ", err)
		return
	}

	listener, err := net.ListenTCP(s.TCPVer, addr)
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

			tcpXconn := NewTCPXConn("1", conn, *s.Router)
			go tcpXconn.Do()
		}
	}()
}

func (s *Server) Serve() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()
	s.Start()
}

func (s *Server) Stop() {

}

func (s *Server) AddRouter(r *intf.XRouter) {
	s.Router = r
}

func New(name string) intf.XServer {
	return &Server{Name: name, TCPVer: "tcp4", Ip: "127.0.0.1", Port: 6069}
}
