/******
** @创建时间 : 2020/3/18 12:47
** @作者 : SongZhiBin
******/
package Utils

import (
	"Songzhibin/GameServer/Giface"
	"fmt"
	"gopkg.in/ini.v1"
)

type Admin struct {
	// server
	TcpServer  Giface.IServer // 当前服务的server对象
	Host       string         // 当前服务器主机监听的ip
	Port       int            // 当前服务器主机监听的端口
	ServerName string         // 当前服务器的名称

	// GameServer
	GVersion       string // 当前程序的版本号
	MaxConn        int    // 当前服务器允许的最大连接数
	MaxMessageSize uint32 // 当前服务器接受数据包的最大值
	MsgChanMaxSize uint32 // 读写管道缓冲最大值
	WorkPollMax    uint32 // 工作池最大容量
	WorkChanMax    uint32 // 任务通道最大长度
	MaxConnection  uint32 // 用户最大连接个数
}

// 全局变量
var AdminObj *Admin

func init() {
	// 初始化AdminObj
	AdminObj = &Admin{
		Host:           "0.0.0.0",
		Port:           18888,
		ServerName:     "Server1",
		GVersion:       "tcp4",
		MaxConn:        1000,
		MaxMessageSize: 4096,
	}
	// 读取配置文件
	cfg, err := ini.Load("/Users/songzhibin/go/src/Songzhibin/GameServer/Game.ini")
	if err != nil {
		fmt.Println(err)
		panic("Read conf file error")
	}
	cfg.Section("Admin").MapTo(AdminObj)
}
