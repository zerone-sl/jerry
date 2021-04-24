package jnet

import (
	"fmt"
	"github.com/zerone/jerry/server/intf"
	"log"
)

type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(r intf.XReq) {
	log.Println("enter router preHandle")
}

func (br *BaseRouter) Handle(r intf.XReq) {
	log.Println("enter router handle")
	cnt, err := r.Conn().Conn().Write(r.Msg())
	if err != nil {
		log.Println("callback err: ", err)
	}
	fmt.Println(cnt)
}

func (br *BaseRouter) PostHandle(r intf.XReq) {
	log.Println("enter router postHandle")
}

func NewRouter() intf.XRouter {
	router := &BaseRouter{}
	return router
}
