package net

import "zinxStudy/szinx/ziface"

type request struct {
	//链接
	conn ziface.IConnection
	//消息
	msg ziface.IMessage
}

//构造方法
func NewRequest(conn ziface.IConnection, msg ziface.IMessage) ziface.IRequest {
	r := &request{
		conn,
		msg,
	}
	return r
}

//得到消息
func (r *request) GetMessage() ziface.IMessage {
	return r.msg
}

//得到当前请求的链接
func (r *request) GetConnection() ziface.IConnection {
	return r.conn
}
