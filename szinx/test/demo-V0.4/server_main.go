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
	request.GetConnection().GetTCPConnection().Write([]byte("\n pre ping\n"))
}
func (r *PingRouter) Handle(request ziface.IRequest){
	request.GetConnection().GetTCPConnection().Write([]byte("\n ping\n"))

	//回显业务
	fmt.Println("【router Handle】 CallBack..")
	data := request.GetData()
	cnt := request.GetDataLen()
	conn := request.GetConnection().GetTCPConnection()

	if _, err := conn.Write(data[:cnt]);err !=nil {
		fmt.Println("write back err ", err)
		return
	}

	return
}
func (r *PingRouter) PostHandle(request ziface.IRequest){
	request.GetConnection().GetTCPConnection().Write([]byte("\n post ping\n"))
}

func main()  {
	s := net.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()
}
