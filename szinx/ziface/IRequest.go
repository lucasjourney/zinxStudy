package ziface

type IRequest interface {
	//得到数据长度方法
	GetDataLen() int
	//得到当前请求的链接
	GetConnection() IConnection
	//得到链接的数据
	GetData() []byte
}
