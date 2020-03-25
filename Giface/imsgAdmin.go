/******
** @创建时间 : 2020/3/19 14:03
** @作者 : SongZhiBin
******/
package Giface

type IMsgAdmin interface {
	// 调用/执行对应Router消息处理方法
	DoMsgHandler(IRequest)
	// 为消息添加的处理逻辑
	AddRouter(uint32, IRouter)
	// Stop 回收工作池
	Stop()
	// 向工作池添加任务
	AppendWork(*IRequest)
}
