package main

import (
	"zinxStudy/szinx/net"
	"zinxStudy/szinx/ziface"
)

type PingRouter struct {
	net.BaseRouter
}

func (r *PingRouter) PreHandle(request ziface.IRequest){

}
func (r *PingRouter) Handle(request ziface.IRequest){
	request.GetConnection().Send(0, []byte("ping..ping..ping"))
	return
}
func (r *PingRouter) PostHandle(request ziface.IRequest){
}

type HelloRouter struct {
	net.BaseRouter
}

func (h *HelloRouter)PreHandle(request ziface.IRequest)  {

}
func (h *HelloRouter)Handle(request ziface.IRequest)  {
	request.GetConnection().Send(0, []byte("hello"))
	return
}
func (h *HelloRouter)PostHandle(request ziface.IRequest)  {

}

func main()  {
	s := net.NewServer()
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &HelloRouter{})
	s.Serve()
}
