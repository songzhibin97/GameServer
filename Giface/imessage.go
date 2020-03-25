/******
** @创建时间 : 2020/3/18 14:58
** @作者 : SongZhiBin
******/
package Giface

type IMessage interface {
	GetMessageId()uint32 // 获取消息id
	GetDataLen()uint32 // 获取消息的长度
	GetData()[]byte // 获取消息内容

	SetMessageID(uint32) // 设置消息id
	SetDataLen(uint32) // 设置消息长度
	SetData([]byte ) // 设置消息内容

}
