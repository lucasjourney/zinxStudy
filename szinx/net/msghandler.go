package net

import (
	"fmt"
	"zinxStudy/szinx/config"
	"zinxStudy/szinx/ziface"
)

//import "zinxStudy/szinx/ziface"

type MsgHandler struct {
	//路由集合
	routers map[uint32]ziface.IRouter
	//负责Worker取任务的消息队列  一个worker对应一个任务队列
	TaskQueues []chan ziface.IRequest
	//worker工作池的worker数量
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		make(map[uint32]ziface.IRouter),
		make([]chan ziface.IRequest, config.GlobalObject.WorkerPoolSize),
		config.GlobalObject.WorkerPoolSize,
	}
}

//启动Worker工作池 (在整个server服务中 只启动一次)
func (mh *MsgHandler) StartWorkerPool() {
	fmt.Println("WorkPool is  started..")

	//根据WorkerPoolSize 创建worker goroutine
	for i := uint32(0); i < config.GlobalObject.WorkerPoolSize; i++ {
		//开启一个workergoroutine

		//1 给当前Worker所绑定消息channel对象 开辟空间  第0个worker 就用第0个Channel
		//给channel 进行开辟空间
		taskQueue := make(chan ziface.IRequest, config.GlobalObject.MaxWorkerTaskLen)
		//2 启动一个Worker，阻塞等待消息从对应的管道中进来
		go mh.startOneWorker(i, taskQueue)
	}

}

//将请求发送给工作池
func (mh *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	//根据链接id来分配给worker
	workerid := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	//分配任务
	mh.TaskQueues[workerid] <- request
}

//一个worker绑定和一个消息队列
func (mh *MsgHandler) startOneWorker(workerID uint32, taskQueue chan ziface.IRequest) {
	fmt.Println("worker ID:", workerID, "is starting...")
	//绑定
	mh.TaskQueues[workerID] = taskQueue
	fmt.Println("taskQueue cap is :", cap(taskQueue))
	for {
		select {
		case req := <-taskQueue:
			mh.DoMsgHandler(req)
		}
	}
}

//添加路由的方法
func (ms *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	//判断msgid 是否存在
	if _, ok := ms.routers[msgID]; ok {
		fmt.Println("该绑定已经存在, msgID：", msgID)
		return
	}
	//添加msgid和router的对应关系
	ms.routers[msgID] = router
	fmt.Println("add router success!")
	return
}

//调度路由的方法
func (ms *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	//取出路由
	router := ms.routers[request.GetMessage().GetMsgId()]
	//执行方法
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}
