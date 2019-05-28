package ziface

/*
##### server模块

###### 属性

ip版本

监听ip（0.0.0.0 全网监控）

监听端口

名称

###### 方法

初始化Server模块的方法

启动服务器

关闭服务器

运行服务器
*/

type IServer interface {
	Serve()
	Close()
	Start()
	AddRouter(msgID uint32, router IRouter)
	GetConnManager() IConnManager

	//注册 创建链接之后的钩子函数
	AddAfterCreateHookFunc(HookFunc)
	//注册 删除链接之前的钩子函数
	AddBeforeDeleteCreateHookFunc(HookFunc)
	//调用 创建链接之后的钩子函数
	CallAfterCreateHookFunc(conn IConnection)
	//调用 删除链接之前的钩子函数
	CallBeforeDeleteCreateHookFunc(conn IConnection)
}

type HookFunc func (conn IConnection)
