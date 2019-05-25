package net

import (
	"errors"
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
		//创建拆包封包对象
		dp := new(DataPack)
		//读取客户端消息的头部
		msgHead := make([]byte, dp.GetHeadLen())

		_, err := io.ReadFull(c.GetTCPConnection(), msgHead)
		if err != nil {
			fmt.Println("read msg head error:", err)
			return
		}
		msg, err := dp.UnPack(msgHead)
		if err != nil {
			fmt.Println("unpack msg error:", err)
			return
		}
		//根据头部，获取数据的长度，读取消息的数据部分
		msgBody := make([]byte, msg.GetMsgLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), msgBody); err != nil {
			fmt.Println("read msg body error:", err)
			return
		}
		msg.SetMsgData(msgBody)
		//将当前一次性得到的对端客户端请求的数据，封装成一个Request
		req := NewRequest(c, msg)

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
func (c *Connection)Send(msgId uint32, msgData []byte) error {
	//判断链接是否结束
	if c.isClosed == true {
		return errors.New("connection is closed....")
	}

	dp := new(DataPack)
	//封装成msg
	binary, err := dp.Pack(NewMessage(msgId, msgData))
	if err != nil {
		fmt.Println("pack err msg id:", msgId)
		return err
	}

	//将binaryMsg发送给对端
	if _, err := c.GetTCPConnection().Write(binary); err != nil {
		fmt.Println("conn write err msg id:",msgId)
		return err
	}
	return nil
}