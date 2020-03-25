/******
** @创建时间 : 2020/3/20 13:29
** @作者 : SongZhiBin
******/
package Gnet

import (
	"Songzhibin/GameServer/Giface"
	"fmt"
	"sync"
)

type ConnectAdmin struct {
	lock  sync.RWMutex
	admin map[uint32]Giface.IConnection
}

// 添加
func (c *ConnectAdmin) Add(connection Giface.IConnection) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.admin[connection.GetConnId()] = connection
	connection.Start()
	fmt.Println("用户已连接", connection.GetConnId(), "ip:", connection.GetRemoteAddr())
}

// 删除
func (c *ConnectAdmin) Remove(connection Giface.IConnection) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.admin, connection.GetConnId())
	fmt.Println("用户已删除", connection.GetConnId(), "ip:", connection.GetRemoteAddr())

}

func (c *ConnectAdmin) GetConnectId(i uint32) (Giface.IConnection, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	v, ok := c.admin[i]
	if !ok {
		return nil, fmt.Errorf("未找到该用户连接")
	}
	return v, nil
}

func (c *ConnectAdmin) RemoveAll() {
	c.lock.Lock()
	defer c.lock.Unlock()
	for k, v := range c.admin {
		v.Stop()
		delete(c.admin, k)
		fmt.Println("以删除并停止 id:", v.GetConnId(), "ip:", v.GetRemoteAddr())
	}
}

func (c *ConnectAdmin) GetLen() uint32 {
	return uint32(len(c.admin))
}
func NewConnectAdmin() *ConnectAdmin {
	return &ConnectAdmin{
		lock:  sync.RWMutex{},
		admin: make(map[uint32]Giface.IConnection),
	}
}
