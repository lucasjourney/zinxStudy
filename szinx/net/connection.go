package net

import (
	"fmt"
	"io"
	"net"
	"zinxStudy/szinx/ziface"
)

type Connection struct {
	//原生套接字 `net.Conn`
	conn *net.TCPConn
	//链接ID `uint32`
	connID uint32
	//当前的conn是否是关闭状态`isClosed bool`
	isClosed bool
	//与当前链接绑定的客户端业务
	router ziface.IRouter
}

//创建一个新的连接
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) ziface.IConnection {
	c := &Connection{
		conn:conn,
		connID:connID,
		isClosed:false,
		router: router,
	}

	return c
}

//针对链接读业务的方法
func (c *Connection) startReader() {
	//从对端读数据
	fmt.Println("Reader go is startin....")
	defer fmt.Println("connID = ", c.connID, "Reader is exit, remote addr is = ", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("recv buf err", err)
			continue
		} else if cnt == 0 {
			fmt.Println("closed by remoteClient , ConnID is :", c.connID)
			break
		}

		req := NewRequest(c, buf, cnt)

		//路由处理业务
		go func() {
			c.router.PreHandle(req)
			c.router.Handle(req)
			c.router.PostHandle(req)
		}()

		//将数据 传递给我们 定义好的Handle Callback方法
		//if err := c.handleAPI(req); err != nil {
		//	fmt.Println("ConnID", c.connID, "Handle is error", err)
		//	break
		//}
	}

}

//	启动链接：当链接模块进行，读写操作
func (c *Connection) Start() {
	fmt.Println("Conn Start（）  ... id = ", c.connID)
	//先进行读业务
	go c.startReader()

	//TODO 进行写业务
}

//	停止链接：关闭套接字/做一些资源回收
func (c *Connection)Stop() {

	//如果已经关闭则返回
	if c.isClosed == true {
		return
	}
	//回收资源
	c.isClosed = true

	c.conn.Close()
}
//	获取链接ID
func (c *Connection)GetConnID() uint32 {
	return c.connID
}
//	获取链接的原生socket套接字
func (c *Connection)GetTCPConnection() *net.TCPConn {
	return c.conn
}
//	查看对端客户端的IP和端口
func (c *Connection)GetRemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}
//	发送数据的方法Send
func (c *Connection)Send(data []byte, cnt int) error {
	if _, err := c.conn.Write(data[:cnt]); err != nil {
		fmt.Println("send buf error")
		return err
	}
	return nil
}