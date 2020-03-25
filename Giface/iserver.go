/******
** @创建时间 : 2020/3/17 12:09
** @作者 : SongZhiBin
******/
package Giface

// server接口层
type IServer interface {
	// 启动方法
	Start()
	// 停止方法
	Stop()
	// 初始化启动
	Server()
	// 添加路由功能 供客户端链接调用使用
	AddRouter(uint32, IRouter)
	// 添加获取当前ConnAdmin的方法
	GetConnAdmin() IConnectAdmin
	// hook对应的方法
	SetOnConnStart(func(IConnection))
	SetOnConnStop(func(IConnection))
	GetOnConnStart(IConnection)
	GetOnConnStop(IConnection)
}
