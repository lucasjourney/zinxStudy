package net

import (
	"errors"
	"fmt"
	"sync"
	"zinxStudy/szinx/ziface"
)

type ConnManager struct {
	//链接集合
	connections map[uint32] ziface.IConnection
	//读写锁
	connLock sync.RWMutex
}

func NewConnManager() ziface.IConnManager {
	return &ConnManager{
		connections:make(map[uint32] ziface.IConnection),
	}
}

//添加链接
func (cm *ConnManager) Add(conn ziface.IConnection) {
	//加锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connections[conn.GetConnID()] = conn

	fmt.Println("Add connid = ", conn.GetConnID(), "to manager succ")
}
//删除链接
func (cm *ConnManager) Remove(connID uint32) {
	//加锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connections, connID)


	fmt.Println("delete connid = ", connID, "from manager succ")

}
//根据链接ID得到链接
func (cm *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	//加读锁
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn, ok := cm.connections[connID]; ok {
		//找到了
		return conn, nil
	} else {
		return nil, errors.New("connection not Found!")
	}
}
//清空链接
func (cm *ConnManager) ClearConn() {
	//加锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	//遍历删除
	for connID, conn := range cm.connections {
		//将全部的conn 关闭
		conn.Stop()

		//删除链接
		delete(cm.connections, connID)
	}

	fmt.Println("Clear All Conections succ! conn num = ", cm.Len())

}
//得到当前服务器链接的总个数
func (cm *ConnManager) Len() uint32 {
	return uint32(len(cm.connections))
}