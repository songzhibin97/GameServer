/******
** @创建时间 : 2020/3/20 13:28
** @作者 : SongZhiBin
******/
package Giface

type IConnectAdmin interface {
	// 添加链接
	Add(connection IConnection)
	// 删除链接
	Remove(connection IConnection)
	// 根据connID获取连接
	GetConnectId(uint32) (IConnection, error)
	// 清空所有
	RemoveAll()
	// 获取连接数
	GetLen() uint32
}
