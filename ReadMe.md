##GameServer 轻量级的TCP服务器框架
===============================

```go
func main() {
	// 1.创建服务
	s := Gnet.NewServer()
	// 注册处理函数
	s.AddRouter(0, new(test))
	s.AddRouter(1, new(test2))
	// 钩子函数
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
```
