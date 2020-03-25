/******
** @创建时间 : 2020/3/18 09:18
** @作者 : SongZhiBin
******/
package Giface

// 路由接口
// 路由的数据都是IRequest请求
type IRouter interface {
	// hook
	// 处理业务之前
	BeforeHandle(IRequest)
	// 处理业务主方法
	NowHandle(IRequest)
	// 处理业务之后
	AfterHandle(IRequest)
}
