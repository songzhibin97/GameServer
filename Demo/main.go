/******
** @创建时间 : 2020/3/17 13:26
** @作者 : SongZhiBin
******/
package main

import (
	"Songzhibin/GameServer/Giface"
	"Songzhibin/GameServer/Gnet"
	"fmt"
)

/*
基于GameServer开发的服务端应用程序
*/

type test struct {
	Gnet.BaseRouter
}

func (t test) BeforeHandle(request Giface.IRequest) {

	fmt.Println("[BeforeHandle]")
	ms := Gnet.InitMessage(request.GetMsgId(), []byte("ping1"))
	back, err := (&Gnet.DataPack{}).Pack(ms)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = request.GetConnection().GetTcpConnection().Write(back)
	if err != nil {
		fmt.Println("error", err)
		return
	}

}
func (t test) NowHandle(request Giface.IRequest) {
	fmt.Println("[NowHandle]")
	ms := Gnet.InitMessage(request.GetMsgId(), []byte("ping2"))
	back, err := (&Gnet.DataPack{}).Pack(ms)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = request.GetConnection().GetTcpConnection().Write(back)
	if err != nil {
		fmt.Println("error", err)
		return
	}

}
func (t test) AfterHandle(request Giface.IRequest) {
	fmt.Println("[AfterHandle]")
	ms := Gnet.InitMessage(request.GetMsgId(), []byte("ping3"))
	back, err := (&Gnet.DataPack{}).Pack(ms)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = request.GetConnection().GetTcpConnection().Write(back)
	if err != nil {
		fmt.Println("error", err)
		return
	}

}

type test2 struct {
	Gnet.BaseRouter
}

func (t test2) BeforeHandle(request Giface.IRequest) {

	fmt.Println("[BeforeHandle]")
	ms := Gnet.InitMessage(request.GetMsgId(), []byte("pong1"))
	back, err := (&Gnet.DataPack{}).Pack(ms)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = request.GetConnection().GetTcpConnection().Write(back)
	if err != nil {
		fmt.Println("error", err)
		return
	}

}
func (t test2) NowHandle(request Giface.IRequest) {
	fmt.Println("[NowHandle]")
	ms := Gnet.InitMessage(request.GetMsgId(), []byte("pong2"))
	back, err := (&Gnet.DataPack{}).Pack(ms)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = request.GetConnection().GetTcpConnection().Write(back)
	if err != nil {
		fmt.Println("error", err)
		return
	}

}
func (t test2) AfterHandle(request Giface.IRequest) {
	fmt.Println("[AfterHandle]")
	ms := Gnet.InitMessage(request.GetMsgId(), []byte("pong3"))
	back, err := (&Gnet.DataPack{}).Pack(ms)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	_, err = request.GetConnection().GetTcpConnection().Write(back)
	if err != nil {
		fmt.Println("error", err)
		return
	}

}

func main() {
	// 1.创建服务
	s := Gnet.NewServer()
	s.AddRouter(0, new(test))
	s.AddRouter(1, new(test2))
	s.SetOnConnStart(func(connection Giface.IConnection) {
		connection.AddAttribute("测试1", "这是测试1数据")
	})
	s.SetOnConnStop(func(connection Giface.IConnection) {
		v, err := connection.GetAttribute("测试1")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(v)
	})
	fmt.Println(s)
	// 启动方法
	s.Server()
}
