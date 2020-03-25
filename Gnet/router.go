/******
** @创建时间 : 2020/3/18 09:18
** @作者 : SongZhiBin
******/
package Gnet

import "Songzhibin/GameServer/Giface"

// 提供基类
// 提供需要继承的基类 基类方法为空
// 好处是如果不需要对应的hook 也不需要强行实现
type BaseRouter struct {
}

func (b *BaseRouter) BeforeHandle(request Giface.IRequest) {

}
func (b *BaseRouter) NowHandle(request Giface.IRequest) {

}
func (b *BaseRouter) AfterHandle(request Giface.IRequest) {

}
