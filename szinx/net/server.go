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

	//链接管理模块
	connManager ziface.IConnManager

	//创建链接之后的钩子函数
	afterCreateHookFunc ziface.HookFunc
	//删除链接之前的钩子函数
	beforeDeleteCreateHookFunc ziface.HookFunc
}

//注册 创建链接之后的钩子函数
func (s *Server)AddAfterCreateHookFunc(hf ziface.HookFunc) {
	s.afterCreateHookFunc = hf
}
//注册 删除链接之前的钩子函数
func (s *Server)AddBeforeDeleteCreateHookFunc(hf ziface.HookFunc) {
	s.beforeDeleteCreateHookFunc = hf
}
//调用 创建链接之后的钩子函数
func (s *Server)CallAfterCreateHookFunc(conn ziface.IConnection) {
	if s.afterCreateHookFunc != nil {
		s.afterCreateHookFunc(conn)
	}
}
//调用 删除链接之前的钩子函数
func (s *Server)CallBeforeDeleteCreateHookFunc(conn ziface.IConnection) {
	if s.beforeDeleteCreateHookFunc != nil {
		s.beforeDeleteCreateHookFunc(conn)
	}
}


func NewServer() ziface.IServer {
	return &Server{
		IPVersion: "tcp4",
		IP:        config.GlobalObject.Host,
		Port:      config.GlobalObject.Port,
		Name:      config.GlobalObject.Name,
		//MaxPackageSize: config.GlobalObject.MaxPackageSize,
		msgHandler:NewMsgHandler(),
		connManager:NewConnManager(),
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


	//开始工作池
	s.msgHandler.StartWorkerPool()
	//3 阻塞等待客户端发送请求，
	go func() {
		for {
			//阻塞等待客户端请求,
			conn, err := listener.AcceptTCP()//只是针对TCP协议
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//最大链接数
			if s.connManager.Len() >= config.GlobalObject.MaxConn {
				//关闭链接
				conn.Close()
				//继续监听
				continue
			} else {
				//创建一个Connection对象
				dealConn := NewConnection(s, conn, cid, s.msgHandler)
				cid++
				//将链接加入管理模块
				s.connManager.Add(dealConn)

				//此时conn就已经和对端客户端连接
				go dealConn.Start()
				//钩子函数
				s.CallAfterCreateHookFunc(dealConn)
			}
		}
	}()
}

func (s *Server) Serve() {
	s.Start()
	//
	select {}
}

func (s *Server) Close() {
	//TODO 关闭所有链接
	s.connManager.ClearConn()
}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter)  {
	s.msgHandler.AddRouter(msgID, router)

}

func (s *Server) GetConnManager() ziface.IConnManager {
	return s.connManager
}