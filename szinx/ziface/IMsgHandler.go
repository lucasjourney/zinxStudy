package ziface

type IMsgHandler interface {
	//添加路由的方法
	AddRouter(msgID uint32, router IRouter)
	//调度路由的方法
	DoMsgHandler(request IRequest)
}
