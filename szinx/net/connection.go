package net

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"zinxStudy/szinx/config"
	"zinxStudy/szinx/ziface"
)

type Connection struct {
	//所属server
	server ziface.IServer
	//原生套接字 `net.Conn`
	conn *net.TCPConn
	//链接ID `uint32`
	connID uint32
	//当前的conn是否是关闭状态`isClosed bool`
	isClosed bool

	//消息管理模块 多路由 消息队列
	msgHandler ziface.IMsgHandler

	//添加一个Reader和Writer通信的Channel
	msgChan chan []byte
	//创建一个Channel 用来通知Writer conn已经关闭
	writerExitChan chan bool

	//链接模块的属性集合
	property map[string] interface{}
	//保护当前的property的锁
	propertyLock sync.RWMutex
}

//创建一个新的连接
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) ziface.IConnection {
	c := &Connection{
		server:         server,
		conn:           conn,
		connID:         connID,
		isClosed:       false,
		msgHandler:     msgHandler,
		msgChan:        make(chan []byte),
		writerExitChan: make(chan bool),
		property:make(map[string]interface{}),
	}

	return c
}

//针对链接读业务的方法
func (c *Connection) startReader() {
	//从对端读数据
	fmt.Println("[Reader Goroutine is startin]....")
	defer fmt.Println("**Reader Goroutine Stop...connID = ", c.connID, "Reader is exit, remote addr is = ", c.GetRemoteAddr().String())
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
		if config.GlobalObject.WorkerPoolSize > 0 {
			c.msgHandler.SendMsgToTaskQueue(req)
		} else {
			go c.msgHandler.DoMsgHandler(req)
		}

	}

}

//写消息的Goroutine 专门负责给客户端发送消息
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is Started]...")
	defer fmt.Println("**Writer Goroutine Stop...")
	//IO多路复用
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.conn.Write(data); err != nil {
				fmt.Println("Send data error", err)
				
				return
			}
		case <-c.writerExitChan:
			//代表reader已经退出了，writer也要退出
			return
		}
	}
}

//	启动链接：当链接模块进行，读写操作
func (c *Connection) Start() {
	fmt.Println("Conn Start（）  ... id = ", c.connID)
	//先进行读业务
	go c.startReader()

	//进行写业务
	go c.StartWriter()

	//TODO 进行写业务
}

//	停止链接：关闭套接字/做一些资源回收
func (c *Connection) Stop() {
	//调用钩子函数
	c.server.CallBeforeDeleteCreateHookFunc(c)

	//如果已经关闭则返回
	if c.isClosed == true {
		return
	}
	//回收资源
	c.isClosed = true
	//通知Writer结束
	c.writerExitChan <- true

	c.conn.Close()

	//将链接在链接管理模块中移除
	c.server.GetConnManager().Remove(c.connID)

	close(c.msgChan)
	close(c.writerExitChan)
}

//	获取链接ID
func (c *Connection) GetConnID() uint32 {
	return c.connID
}

//	获取链接的原生socket套接字
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.conn
}

//	查看对端客户端的IP和端口
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

//	发送数据的方法Send
func (c *Connection) Send(msgId uint32, msgData []byte) error {
	//判断链接是否结束
	if c.isClosed == true {
		return errors.New("connection is closed....")
	}

	dp := new(DataPack)
	//封装成msg
	binaryMsg, err := dp.Pack(NewMessage(msgId, msgData))
	if err != nil {
		fmt.Println("pack err msg id:", msgId)
		return err
	}

	//将要打包好的二进制数据发送给channel，让writer去写
	c.msgChan <- binaryMsg

	return nil
}

//  设置属性
func (c *Connection)SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}
//  获取属性
func (c *Connection)GetProperty(key string)(interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("error:connection property do not have property:" + key)
	}
}
//  删除属性
func (c *Connection)RemoveProperty(key string) {
	delete(c.property, key)
}