/******
** @创建时间 : 2020/3/17 15:58
** @作者 : SongZhiBin
******/
package main

import (
	"Songzhibin/GameServer/Gnet"
	"fmt"
	"net"
	"time"
)

func main() {
	for i := 0; i < 2; i++ {
		go func(i int) {
			conn, err := net.Dial("tcp4", "127.0.0.1:18887")
			if err != nil {
				fmt.Println(err)
				return
			}
			for {
				fmt.Println(i, "xxxxxxxxxxxxxxxxxxxxxxx")
				s := fmt.Sprint("hello")
				ms := Gnet.InitMessage(uint32(i), []byte(s))
				back, err := (&Gnet.DataPack{}).Pack(ms)
				if err != nil {
					fmt.Println(err)
					return
				}
				_, err = conn.Write(back)
				if err != nil {
					fmt.Println("error", err)
					return
				}
				for {
					ms, err := Gnet.GetMsg(conn)
					if err != nil {
						fmt.Println(err)
						return
					}
					fmt.Println("ms: id:", (*ms).GetMessageId(), "data:", string((*ms).GetData()))
				}
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
