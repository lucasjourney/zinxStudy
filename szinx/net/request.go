package net

import "zinxStudy/szinx/ziface"

type request struct {
	//链接
	conn ziface.IConnection
	//数据
	data []byte
	//数据长度
	len int
}

//构造方法
func NewRequest(conn ziface.IConnection, data []byte, len int) ziface.IRequest {
	r := &request{
		conn,
		data,
		len,
	}
	return r
}

//得到数据长度方法
func (r *request)GetDataLen() int {
	return r.len
}
//得到当前请求的链接
func (r *request)GetConnection() ziface.IConnection {
	return r.conn
}
//得到链接的数据
func (r *request)GetData()  []byte {
	return r.data
}