package ziface

//链接管理接口
type IConnManager interface {
	//添加链接
	Add(conn IConnection)
	//删除链接
	Remove(connID uint32)
	//根据链接ID得到链接
	Get(connID uint32) (IConnection, error)
	//清空链接
	ClearConn()
	//得到当前服务器链接的总个数
	Len() uint32
}
