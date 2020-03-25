/******
** @创建时间 : 2020/3/17 17:24
** @作者 : SongZhiBin
******/
package Gnet

import (
	"Songzhibin/GameServer/Giface"
	"Songzhibin/GameServer/Utils"
	"fmt"
	"io"
	"net"
	"sync"
)

//
type Connection struct {
	// socket句柄
	conn *net.TCPConn
	// 链接的id
	connId uint32
	// 当前的链接状态
	isClosed bool
	// 当前链接所绑定的处理业务方法
	//HandleApi Giface.HandleFunc // 新增router后剔除
	// 通知是否退出channel
	exitChannel chan bool
	// 新增Router成员
	//Router Giface.IRouter v0.6进行修改
	routers Giface.IMsgAdmin
	// v0.7新增 channel 用于读写通讯 无缓冲通道
	msgChan chan []byte
	// v0.9新增 server
	server Giface.IServer
	// v1.0新增 管理属性
	attributeAdmin map[string]interface{}
	// 属性读写锁
	attributeLock sync.RWMutex
}

func (c *Connection) startReader() {
	fmt.Println("进入start Read")
	fmt.Println("[Read Start]Read Goroutine start id:", c.connId)
	defer fmt.Println("[Read Over] Read Over id:", c.connId, "remote addr is : ", c.GetRemoteAddr().String())
	defer c.Stop()
	// v0.5废弃
	//// 读取客户端数据到buf中
	//buf := make([]byte, Utils.AdminObj.MaxMessageSize)
	//num, err := c.Conn.Read(buf)
	//if err != nil {
	//	if err != io.EOF {
	//		return
	//	}
	//	fmt.Println("读取失败 err", err, "id:", c.ConnId)
	//	break
	//}
	//// 得到当前conn读取数据的封装 Request
	//req := NewRequest(c, buf[:num])
	//// 调用当前链接绑定的 从路由中找到注册绑定的conn对应的router调用
	//go func(req Giface.IRequest) {
	//	c.Router.BeforeHandle(req)
	//	c.Router.NowHandle(req)
	//	c.Router.AfterHandle(req)
	//}(req)

	// v0.5 继承datapack后

	for {
		// 封装方法
		ms, err := GetMsg(c.GetTcpConnection())
		fmt.Println("读取数据", c.connId)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("[Reconnect 重新连接]", err, "id:", c.connId)
			continue
		}
		fmt.Println("准备添加任务 id:", c.connId)
		// 优化 如果消息长度大于0才进行处理
		if (*ms).GetDataLen() > 0 {
			// 得到当前conn读取数据的封装 Request
			req := NewRequest(c, *ms)
			// MsgAdmin 根据对应的id运行对应的方法
			fmt.Println("添加任务 id:", c.connId)
			c.routers.AppendWork(&req)
		} else {
			fmt.Println("[data 0 continue] id:", c.connId)
		}
	}

}

// v0.7新增 write协程 专门将消息发送到客户端模块
func (c *Connection) startWrite() {
	fmt.Println("[Write Start] Write Goroutine start id:", c.connId)
	defer fmt.Println("[Write Over] Write Over id:", c.connId, "remote addr is : ", c.GetRemoteAddr().String())
	select {
	case data := <-c.msgChan:
		// 如果有消息进行发送
		if _, err := c.GetTcpConnection().Write(data); err != nil {
			fmt.Println("error", err)
			return
		}
	case <-c.exitChannel:
		// 检测到退出进行退出
		return
	}
}

func (c *Connection) Start() {
	go c.startReader()
	go c.startWrite()
	fmt.Println("Conn Start id:", c.connId)

	// v0.9新增hook入口

}
func (c *Connection) Stop() {
	fmt.Println("Conn Stop id:", c.connId)
	if c.isClosed {
		return
	}
	c.server.GetOnConnStop(c)
	c.isClosed = true
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil { //判断是否报错 如果==nil没有抛错
			fmt.Println("recover err :", err)
			return
		}
	}()

	close(c.exitChannel)
	close(c.msgChan)
	// 通知write协程退出
	c.server.GetConnAdmin().Remove(c)

}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.conn
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) GetConnId() uint32 {
	return c.connId
}

func (c *Connection) Send(message Giface.IMessage) error {
	// 提供一个封装的方法 将message 封装为 []byte 发送
	// 判断是否通道已经关闭
	if c.isClosed {
		return fmt.Errorf("Connect is close\n")
	}
	back, err := (&DataPack{}).Pack(message)
	if err != nil {
		return err
	}
	// v0.7更改 不进行直接写入 放入channel中
	//_, err = c.GetTcpConnection().Write(back)
	//if err != nil {
	//	return err
	//}
	c.msgChan <- back
	return nil
}

// v1.0新增 增删查属性
func (c *Connection) AddAttribute(name string, attr interface{}) {
	c.attributeLock.Lock()
	defer c.attributeLock.Unlock()
	c.attributeAdmin[name] = attr
}
func (c *Connection) DelAttribute(name string) {
	c.attributeLock.Lock()
	defer c.attributeLock.Unlock()
	delete(c.attributeAdmin, name)
}
func (c *Connection) GetAttribute(name string) (interface{}, error) {
	c.attributeLock.RLock()
	defer c.attributeLock.RUnlock()
	v, ok := c.attributeAdmin[name]
	if !ok {
		return nil, fmt.Errorf("无此属性 name:%s", name)
	}
	return v, nil

}

func GetMsg(c net.Conn) (*Giface.IMessage, error) {
	headBuf := make([]byte, (&DataPack{}).GetHeadLen())
	// 第一次读 获取head部分
	_, err := io.ReadFull(c, headBuf)
	if err == io.EOF {
		return nil, err
	}
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	// 解析head
	ms, err := (&DataPack{}).UnPack(headBuf)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if ms.GetDataLen() > 0 {
		// 第二次读取
		data := make([]byte, ms.GetDataLen())
		_, err := io.ReadFull(c, data)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		ms.SetData(data)
	}
	return &ms, nil
}

func NewConnection(server Giface.IServer, conn *net.TCPConn, connId uint32) *Connection {

	res := &Connection{
		conn:     conn,
		connId:   connId,
		isClosed: false,
		//HandleApi:   handleApi, 新增router后handleApi进行剔除
		exitChannel: make(chan bool, 1),
		//Router:      router,  v0.6中移出
		routers: NewMsgAdmin(),
		// v0.7读写分离新增
		msgChan:        make(chan []byte, Utils.AdminObj.MsgChanMaxSize),
		server:         server,
		attributeAdmin: make(map[string]interface{}),
	}
	// v0.9新增启动
	res.server.GetOnConnStart(res)
	res.server.GetConnAdmin().Add(res)
	res.Start()

	return res
}
