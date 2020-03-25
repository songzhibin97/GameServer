/******
** @创建时间 : 2020/3/17 17:24
** @作者 : SongZhiBin
******/
package Giface

import "net"

// connection接口层
type IConnection interface {
	// 启动链接
	Start()
	// 关闭连接
	Stop()
	// 获取当前链接的绑定socket 句柄
	GetTcpConnection() *net.TCPConn
	// 获取远程客户端Tcp的状态 ip:port
	GetRemoteAddr() net.Addr
	// 获取当前链接模块的id
	GetConnId() uint32
	// 发送数据给当前的远程客户端
	Send(message IMessage) error
	// v1.0新增 管理连接属性
	AddAttribute(string, interface{})
	DelAttribute(string)
	GetAttribute(string) (interface{}, error)
}

// 实现业务函数
type HandleFunc func(*net.TCPConn, []byte, int) error
