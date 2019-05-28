package main

import (
	"fmt"
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
	s.AddAfterCreateHookFunc(func(conn ziface.IConnection) {
		fmt.Println("===afterHookFunc:")
		//链接一旦创建成功 给用户返回一个消息
		if err := conn.Send(202, []byte("Hello welcome to zinx...")); err !=nil {
			fmt.Println(err)
		}
		//当用户一旦链接创建成功， 给链接绑定一些属性
		fmt.Println("Set conn property...")
		conn.SetProperty("Name", "Go3")
		conn.SetProperty("address", "TDB...")
		conn.SetProperty("time", "2019-12-12")
	})
	s.AddBeforeDeleteCreateHookFunc(func(conn ziface.IConnection) {
		fmt.Println("===beforeHookFunc:")
		fmt.Println("Conn id ", conn.GetConnID(), "is Lost!.")

		fmt.Println("Get Conn Property...")
		//获取conn Name
		if name, err := conn.GetProperty("Name"); err == nil {
			fmt.Println("Name =", name)
		}
		//获取conn address
		if address, err := conn.GetProperty("address"); err == nil {
			fmt.Println("address =", address)
		}
		//获取conn time
		if time, err := conn.GetProperty("time"); err == nil {
			fmt.Println("address =", time)
		}
	})
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &HelloRouter{})
	s.Serve()
}
