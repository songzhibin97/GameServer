/******
** @创建时间 : 2020/3/17 21:33
** @作者 : SongZhiBin
******/
package Gnet

import (
	"Songzhibin/GameServer/Giface"
)

type Request struct {
	// 句柄
	conn Giface.IConnection
	// 数据data
	//Data []byte v0.5 继承datapack后更换
	msg Giface.IMessage
}

func (r *Request) GetConnection() Giface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}
func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMessageId()
}

func NewRequest(conn Giface.IConnection, message Giface.IMessage) Giface.IRequest {
	return &Request{
		conn: conn,
		msg:  message,
	}
}
