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
}

//配置一个全局变量
var GlobalObject *GlobalObj

func init()  {
	GlobalObject = &GlobalObj{
		//默认值
		Host:"127.0.0.1",
		Port:9999,
		Name:"zinx app",
		Version:"v0.4",
		MaxPackageSize:4096,
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
