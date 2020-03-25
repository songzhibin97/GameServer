/******
** @创建时间 : 2020/3/19 14:03
** @作者 : SongZhiBin
******/
package Gnet

import (
	"Songzhibin/GameServer/Giface"
	"Songzhibin/GameServer/Utils"
	"context"
	"fmt"
)

var LocalMsgAdmin *MsgAdmin

type MsgAdmin struct {
	admins   map[uint32]Giface.IRouter
	workChan chan *Giface.IRequest
	cancel   context.CancelFunc
}

func (m *MsgAdmin) DoMsgHandler(request Giface.IRequest) {
	// 1. request中找到msgID
	router, ok := m.admins[request.GetMsgId()]
	if !ok {
		fmt.Println("id:", request.GetMsgId(), "找不到对应路由,调度失败")
		return
	}
	// 2. 根据msgId在MsgAdmin中找到对应的router进行运行
	router.AfterHandle(request)
	router.NowHandle(request)
	router.BeforeHandle(request)
}

func (m *MsgAdmin) AddRouter(i uint32, request Giface.IRouter) {
	// 判断是否注册
	if _, ok := m.admins[i]; ok {
		fmt.Println("id:", i, " 已经注册")
		return
	}
	m.admins[i] = request
	fmt.Println("id:", i, " 注册成功")
}

func NewMsgAdmin() *MsgAdmin {
	if LocalMsgAdmin != nil {
		return LocalMsgAdmin
	}
	LocalMsgAdmin = &MsgAdmin{admins: make(map[uint32]Giface.IRouter),
		workChan: make(chan *Giface.IRequest, Utils.AdminObj.WorkChanMax)}
	return LocalMsgAdmin
}

// worker
func (m *MsgAdmin) worker(i int, ctx context.Context, jobs <-chan *Giface.IRequest) {

	for {
		select {
		case req := <-jobs:
			fmt.Println(i, "接受任务", (*req).GetMsgId())
			m.DoMsgHandler(*req)
		case <-ctx.Done():
			return
		}
	}
}

func (m *MsgAdmin) Stop() {
	m.cancel()
}

// 新增方法 向任务池发送任务
func (m *MsgAdmin) AppendWork(request *Giface.IRequest) {
	m.workChan <- request
}

// 初始化方法 需要在server中调用
func MsgInit() {
	ctx, cancel := context.WithCancel(context.Background())
	LocalMsgAdmin = &MsgAdmin{admins: make(map[uint32]Giface.IRouter),
		workChan: make(chan *Giface.IRequest, Utils.AdminObj.WorkChanMax),
		cancel:   cancel}
	fmt.Println(len(LocalMsgAdmin.workChan),cap(LocalMsgAdmin.workChan))
	for i := 0; i < int(Utils.AdminObj.WorkPollMax); i++ {
		go LocalMsgAdmin.worker(i, ctx, LocalMsgAdmin.workChan)
	}
}
