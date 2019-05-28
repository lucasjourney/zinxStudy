package config

import (
	"encoding/json"
	"io/ioutil"
)

type GlobalObj struct {
	Host string
	Port int
	Name string

	Version string
	MaxPackageSize uint32

	//工作池worker数量
	WorkerPoolSize uint32
	//消息队列长度
	MaxWorkerTaskLen uint32
	//最大链接数
	MaxConn uint32
}

//配置一个全局变量
var GlobalObject *GlobalObj

func init()  {
	GlobalObject = &GlobalObj{
		//默认值
		Host:"127.0.0.1",
		Port:7777,
		Name:"zinx app",
		Version:"v0.5",
		MaxPackageSize:4096,
		WorkerPoolSize:4,
		MaxWorkerTaskLen:8,
		MaxConn:4,
	}

	GlobalObject.loadConfig()
}

func (g *GlobalObj) loadConfig()  {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, GlobalObject)
	if err != nil {
		panic(err)
	}
}
