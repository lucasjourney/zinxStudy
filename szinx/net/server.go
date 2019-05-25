package net

import (
	"fmt"
	"zinxStudy/szinx/config"

	//"io"
	"net"
	"zinxStudy/szinx/ziface"
)

type Server struct {
	IPVersion string
	IP        string
	Port      int
	Name      string
	//router	  ziface.IRouter
	//MaxPackageSize uint32

	//消息管理模块 多路由
	msgHandler ziface.IMsgHandler
}


func NewServer() ziface.IServer {
	return &Server{
		IPVersion: "tcp4",
		IP:        config.GlobalObject.Host,
		Port:      config.GlobalObject.Port,
		Name:      config.GlobalObject.Name,
		//MaxPackageSize: config.GlobalObject.MaxPackageSize,
		msgHandler:NewMsgHandler(),
	}
}

func (s *Server) Start() {
	fmt.Printf("[start] Server Linstenner at IP :%s, Port :%d, is starting..\n", s.IP, s.Port)
	//1.创建套接字
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr err:", err)
		return
	}
	//2.监听服务器地址
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen tcp err:", err)
		return
	}

	//生成id的累加器
	var cid uint32
	cid = 0

	//3 阻塞等待客户端发送请求，
	go func() {
		for {
			//阻塞等待客户端请求,
			conn, err := listener.AcceptTCP()//只是针对TCP协议
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//创建一个Connection对象
			dealConn := NewConnection(conn, cid, s.msgHandler)
			cid++


			//此时conn就已经和对端客户端连接
			go dealConn.Start()
		}
	}()
}

func (s *Server) Serve() {
	s.Start()
	//
	select {}
}

func (s *Server) Close() {
	//TODO
}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter)  {
	s.msgHandler.AddRouter(msgID, router)
	fmt.Println("add router success!")
}
