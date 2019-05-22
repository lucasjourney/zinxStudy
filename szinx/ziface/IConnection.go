package ziface

import "net"

type IConnection interface {
	//	启动链接：当链接模块进行，读写操作
	Start()
	//	停止链接：关闭套接字/做一些资源回收
	Stop()
	//	获取链接ID
	GetConnID() uint32
	//	获取链接的原生socket套接字
	GetTCPConnection() *net.TCPConn
	//	查看对端客户端的IP和端口
	GetRemoteAddr() net.Addr
	//	发送数据的方法Send
	Send(data []byte) error
}

//业务处理方法 抽象定义
type HandleFunc func(*net.TCPConn, []byte, int) error
