/**
 zinx v0.1 应用
*/
package main

import (
	"zinxStudy/szinx/net"
)

func main() {
	//创建一个zinx server对象
	s := net.NewServer("zinx v0.1")

	//注册一些自定义的业务
	//s.AddRouter(1, &dsad)

	//让server对象 启动服务
	s.Serve()

	return
}
