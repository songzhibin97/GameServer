/******
** @创建时间 : 2020/3/18 15:59
** @作者 : SongZhiBin
******/
package Gnet

import (
	"Songzhibin/GameServer/Giface"
	"Songzhibin/GameServer/Utils"
	"bytes"
	"encoding/binary"
	"fmt"
)

type DataPack struct {
}

func (d *DataPack) GetHeadLen() uint32 {
	// packData
	// head:8byte id 8byte data 16: byte
	return 8
}

// 封包
func (d *DataPack) Pack(msg Giface.IMessage) ([]byte, error) {
	// head(len) -> 8byte
	// id -> 8byte
	// data -> 16:len(data)

	// 创建[]data
	buf := new(bytes.Buffer)
	// 1.将dataLen写入 buf中
	if err := binary.Write(buf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	// 2.将id写入 buf中
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMessageId()); err != nil {
		return nil, err
	}
	// 3.将data数据写入 buf中
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (d *DataPack) UnPack(data []byte) (Giface.IMessage, error) {
	// 这里的data是获取请求头的数据 已经解析过得 data[:d.GetHeadLen()]
	// byte转化为reader对象
	buf := bytes.NewReader(data)
	msg := NewMessage()
	if err := binary.Read(buf, binary.LittleEndian, &msg.dataLen); err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	if err := binary.Read(buf, binary.LittleEndian, &msg.id); err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	// 判断配置是否已经超出最大长度
	if Utils.AdminObj.MaxMessageSize < msg.dataLen {
		fmt.Println(Utils.AdminObj.MaxMessageSize, msg.dataLen)
		return nil, fmt.Errorf("exceed maxsize")
	}
	return msg, nil
}

// 实例化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}
