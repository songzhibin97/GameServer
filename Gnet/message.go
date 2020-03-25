/******
** @创建时间 : 2020/3/18 14:57
** @作者 : SongZhiBin
******/
package Gnet

type Message struct {
	// 消息的id
	id uint32
	// 消息的长度
	dataLen uint32
	// 消息的内容
	data []byte
}

func (m *Message) GetMessageId() uint32 {
	return m.id
}
func (m *Message) GetDataLen() uint32 {
	return m.dataLen
}
func (m *Message) GetData() []byte {
	return m.data
}
func (m *Message) SetMessageID(id uint32) {
	m.id = id
}
func (m *Message) SetDataLen(l uint32) {
	m.dataLen = l
}
func (m *Message) SetData(data []byte) {
	m.data = data
}
func NewMessage() *Message {
	return &Message{}
}

// 进行二次封装
func InitMessage(id uint32, data []byte) *Message {
	msg := NewMessage()
	msg.SetMessageID(id)
	msg.SetDataLen(uint32(len(data)))
	msg.SetData(data)
	return msg
}
