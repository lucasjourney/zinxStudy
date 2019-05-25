package ziface

type IDataPack interface {
	//获取头部长度
	GetHeadLen() uint32
	//封包 打包成|datalen|dataID|data|
	Pack(msg IMessage) ([]byte, error)
	//拆包
	UnPack([]byte) (IMessage, error)
}
