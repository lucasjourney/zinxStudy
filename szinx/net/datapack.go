package net

import (
	"bytes"
	"encoding/binary"
	"zinxStudy/szinx/ziface"
)

type DataPack struct {

}

func NewDataPack() *DataPack {
	return new(DataPack)
}

//获取头部长度
func (dp *DataPack) GetHeadLen() uint32 {
	//uint32 + uint32 长度为8
	return 8
}

//封包 打包成|datalen|dataID|data|
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放二进制的字节缓冲
	buff := bytes.NewBuffer([]byte{})
	//将datalen 写进buffer中
	if err := binary.Write(buff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	//将 dataID写进buffer中
	if err := binary.Write(buff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	//将 data写进buffer中
	if err := binary.Write(buff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}
	//返回这个缓冲
	return buff.Bytes(), nil
}

//拆包 将长度、ID、数据拆包到Message结构体中，
//解包的时候 是分2次解压，第一次读取固定的长度8字节， 第二次是根据len 再次进行read
func (dp *DataPack) UnPack(binarydata []byte) (ziface.IMessage, error) {
	//
	msgHead := new(Message)
	//创建一个 读取二进制数据流的io.Reader
	databuff := bytes.NewReader(binarydata)
	//将二进制流 先读datalen 放在msg的DataLen属性中
	if err := binary.Read(databuff, binary.LittleEndian, &msgHead.Datalen); err != nil {
		return nil, err
	}
	//将二进制流的  DataID 方在Msg的DataID属性中
	if err := binary.Read(databuff, binary.LittleEndian, &msgHead.Id); err != nil {
		return nil, err
	}

	return msgHead, nil
}
