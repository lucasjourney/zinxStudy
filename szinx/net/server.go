package net

import (
	"fmt"
	"io"
	"net"
	"szinx/ziface"
)

type Server struct {
	IPVersion string
	IP        string
	Port      int
	Name      string
}

func NewServer(name string) ziface.IServer {
	return &Server{
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Name:      name,
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

	go func() {
		for {
			//3.阻塞等待客户端发送请求
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("listen accept err:", err)
				return
			}

			go func() {
				for {
					//4.客户端有数据请求，处理客户端业务（读写）
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil && err != io.EOF {
						fmt.Println("conn read err:", err)
						return
					} else if cnt == 0 {
						fmt.Println("read finished")
						return
					}

					fmt.Printf("recv client buf %s, cnt = %d\n", buf, cnt)

					//回显功能 （业务）
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
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
