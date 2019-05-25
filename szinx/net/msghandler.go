package net

import (
	"fmt"
	"zinxStudy/szinx/ziface"
)

//import "zinxStudy/szinx/ziface"

type MsgHandler struct {
	//路由集合
	routers map[uint32] ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		make(map[uint32]ziface.IRouter),
	}
}

//添加路由的方法
func (ms *MsgHandler)AddRouter(msgID uint32, router ziface.IRouter) {
	//判断msgid 是否存在
	if _, ok := ms.routers[msgID]; ok {
		fmt.Println("该绑定已经存在, msgID：", msgID)
		return
	}
	//添加msgid和router的对应关系
	ms.routers[msgID] = router
	return
}
//调度路由的方法
func (ms *MsgHandler)DoMsgHandler(request ziface.IRequest) {
	//取出路由
	router := ms.routers[request.GetMessage().GetMsgId()]
	//执行方法
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}