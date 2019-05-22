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
}
