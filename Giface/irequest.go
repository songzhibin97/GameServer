/******
** @创建时间 : 2020/3/17 21:33
** @作者 : SongZhiBin
******/
package Giface

// 实际上是把客户端请求的链接信息和请求的数据包装到一个Request中
type IRequest interface {
	// 获取当前链接
	GetConnection() IConnection

	// 获取请求的数据
	GetData() []byte

	// 获取msg id
	GetMsgId() uint32
}
