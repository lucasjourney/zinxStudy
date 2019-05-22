package net

import (
	"net"
	"szinx/ziface"
)

type Connection struct {
	//原生套接字 `net.Conn`
	conn net.Conn
	//链接ID `uint32`
	connID uint32
	//当前的conn是否是关闭状态`isClosed bool`
	isClosed bool
	//与当前链接绑定的客户端业务
	handleAPI ziface.HandleFunc
}

//创建一个新的连接
func NewConnection(conn net.Conn, connID uint32, isClosed bool,
	handleAPI ziface.HandleFunc) ziface.IConnection {
	c := &Connection{
		conn:conn,
		connID:connID,
		isClosed:isClosed,
		handleAPI:handleAPI,
	}

	return c
}

//	启动链接：当链接模块进行，读写操作
func (conn *Connection)Start() {

}
//	停止链接：关闭套接字/做一些资源回收
func (conn *Connection)Stop() {

}
//	获取链接ID
func (conn *Connection)GetConnID() uint32 {
	return 0
}
//	获取链接的原生socket套接字
func (conn *Connection)GetTCPConnection() *net.TCPConn {
	return nil
}
//	查看对端客户端的IP和端口
func (conn *Connection)GetRemoteAddr() net.Addr {
	return nil
}
//	发送数据的方法Send
func (conn *Connection)Send(data []byte) error {
	return nil
}