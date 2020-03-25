/******
** @创建时间 : 2020/3/17 12:11
** @作者 : SongZhiBin
******/
package Gnet

import (
	"Songzhibin/GameServer/Giface"
	"Songzhibin/GameServer/Utils"
	"fmt"
	"net"
	"sync/atomic"
)

var LOGO = `                                        
	=======		 	==			 ========
	==    ===	 	==			==	==
	==     ===	 				==	==
	==    ===	 	==			==	==
	=======		 	==			==	==
	==    ====	 	==			==	==	
	==     === 	 	==			==	==
	==    ===	 	==			==	==
	=======		 	==			==	==
											
`

// 定义Server结构体 实现IServer方法
type Server struct {
	// 服务器名称
	serverName string
	// 绑定IP版本
	iPVersion string
	// 服务IP
	serverIP string
	// 服务端口
	serverPort int
	// 增加Router成员
	//Router Giface.IRouter v0.6进行更换
	routers Giface.IMsgAdmin
	// v0.9新增 server连接管理器
	connAdmin Giface.IConnectAdmin
	// v0.9新增 hook函数
	onConnStart func(connection Giface.IConnection)
	onConnStop  func(connection Giface.IConnection)
}

func (s *Server) Start() {
	// 单体服务的实现
	fmt.Printf("Server Start Listen at IP:%s, Port:%d\n", s.serverIP, s.serverPort)
	// 1.获取一个TCP的address
	addr, err := net.ResolveTCPAddr(s.iPVersion, fmt.Sprintf("%s:%d", s.serverIP, s.serverPort))
	if err != nil {
		fmt.Printf("addr error:%s\n", err)
		return
	}
	// 2.尝试监听服务的地址
	listen, err := net.ListenTCP(s.iPVersion, addr)
	if err != nil {
		fmt.Printf("listen err:%s\n", err)
		return
	}
	fmt.Printf("Server Success Listen at IP:%s, Port:%d\n", s.serverIP, s.serverPort)
	// 3.阻塞的等待客户端连接 处理客户端链接业务
	var id uint32

	for {
		// 如果有用户连接 创建goroutine进行异步处理
		tcpConn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("connect error", err)
			continue
		}
		atomic.AddUint32(&id, 1)
		go func(id uint32, tcpConn *net.TCPConn) {
			// conn客户端句柄 进行读写操作(业务层 可以实现对应业务)
			if s.connAdmin.GetLen() > Utils.AdminObj.MaxConnection {
				fmt.Println("[max connect]连接已超限,当前链接数量", s.connAdmin.GetLen())
				// 超过最大连接数 对用户连接发送连接失败的信息 并关闭此连接
				(&DataPack{}).Pack(InitMessage(0, []byte("连接超过最大值,连接被迫关闭")))
				// 关闭连接
				tcpConn.Close()
				return
			}
			fmt.Println("创建Connection 读写: id", atomic.LoadUint32(&id))
			NewConnection(s, tcpConn, atomic.LoadUint32(&id))
			// v0.9移出 NewConnection Add 中启动
			//// 启动
			//go dealConn.Start()
		}(id, tcpConn)
	}
}

func (s *Server) GetConnAdmin() Giface.IConnectAdmin {
	return s.connAdmin
}
func (s *Server) Stop() {
	// todo 停止回收响应资源
	// 关闭当前所有用户的链接
	s.routers.Stop()
	s.connAdmin.RemoveAll()

}
func (s *Server) Server() {
	// 启动运行服务 调用内置的start方法
	fmt.Println(LOGO)
	s.Start()
	// todo 做一些其他业务
}
func (s *Server) AddRouter(i uint32, r Giface.IRouter) {
	s.routers.AddRouter(i, r)
	fmt.Println("[router] 新增路由方法成功")
}

// 注册 hook函数的方法
func (s *Server) SetOnConnStart(f func(connection Giface.IConnection)) {
	fmt.Println("hook start add")
	s.onConnStart = f
}
func (s *Server) SetOnConnStop(f func(connection Giface.IConnection)) {
	fmt.Println("hook end add")
	s.onConnStop = f
}

// 调用 hook函数的方法
func (s *Server) GetOnConnStart(c Giface.IConnection) {
	if s.onConnStart != nil {
		fmt.Println("[server hook start]")
		s.onConnStart(c)
	}

}
func (s *Server) GetOnConnStop(c Giface.IConnection) {
	if s.onConnStop != nil {
		fmt.Println("[server hook over]")
		s.onConnStop(c)
	}
}

func NewServer() Giface.IServer {
	// 初始化线程池
	MsgInit()
	// 单例
	if Utils.AdminObj.TcpServer != nil {
		return Utils.AdminObj.TcpServer
	}
	Utils.AdminObj.TcpServer = &Server{
		serverName: Utils.AdminObj.ServerName,
		iPVersion:  Utils.AdminObj.GVersion,
		serverIP:   Utils.AdminObj.Host,
		serverPort: Utils.AdminObj.Port,
		routers:    NewMsgAdmin(),
		connAdmin:  NewConnectAdmin(),
	}
	return Utils.AdminObj.TcpServer
}
