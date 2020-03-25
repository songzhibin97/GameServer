/******
** @创建时间 : 2020/3/18 20:57
** @作者 : SongZhiBin
******/
package Gnet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestDataPack(t *testing.T) {
	listen, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println(err)
		return
	}
	go func(listen net.Listener) {
		for {
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println(err)
				return
			}
			go func(c net.Conn) {
				for {
					headBuf := make([]byte, (&DataPack{}).GetHeadLen())
					// 第一次读 获取head部分
					_, err := io.ReadFull(conn, headBuf)
					if err != nil {
						fmt.Println(err)
						return
					}
					// 解析head
					ms, err := (&DataPack{}).UnPack(headBuf)
					if err != nil {
						fmt.Println(err)
						return
					}
					if ms.GetDataLen() > 0 {
						// 第二次读取
						data := make([]byte, ms.GetDataLen())
						_, err := io.ReadFull(conn, data)
						if err != nil {
							fmt.Println(err)
							return
						}
						ms.SetData(data)
						fmt.Println(ms.GetMessageId(), ms.GetDataLen(), string(ms.GetData()))
					}
				}
			}(conn)
		}
	}(listen)

	go func() {
		fmt.Println(111)
		listen, err := net.Dial("tcp4", "127.0.0.1:7777")
		if err != nil {
			fmt.Println(err)
			return
		}
		mes := NewMessage()
		mes.SetMessageID(2)
		mes.SetDataLen(7)
		mes.SetData([]byte{'h', 'e', 'l', 'l', 'o', 'w', '!'})

		back, err := (&DataPack{}).Pack(mes)
		if err != nil {
			fmt.Println(err)
		}
		listen.Write(back)
	}()
	go func() {
		fmt.Println(222)
		listen, err := net.Dial("tcp4", "127.0.0.1:7777")
		if err != nil {
			fmt.Println(err)
			return
		}
		mes := NewMessage()
		mes.SetMessageID(10)
		mes.SetDataLen(6)
		mes.SetData([]byte{'h', 'e', 'l', 'l', 'o', 'w'})

		back, err := (&DataPack{}).Pack(mes)
		if err != nil {
			fmt.Println(err)
		}
		listen.Write(back)
	}()
	time.Sleep(time.Minute)
}
