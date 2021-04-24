package intf

type XRouter interface {
	PreHandle(r XReq)

	Handle(r XReq)

	PostHandle(r XReq)
}
