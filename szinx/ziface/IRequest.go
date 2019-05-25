package ziface

type IRequest interface {

	//得到当前请求的链接
	GetConnection() IConnection
	//得到消息
	GetMessage() IMessage
}
