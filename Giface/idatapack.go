/******
** @创建时间 : 2020/3/18 15:59
** @作者 : SongZhiBin
******/
package Giface

type IDataPack interface {
	// 获取包头长度方法
	GetHeadLen() uint32

	// 封包方法
	Pack(IMessage) ([]byte, error)
	// 拆包方法 主要是根据固定头长度获取数据长度和id
	UnPack([]byte) (IMessage,error)
}
