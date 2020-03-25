/******
** @创建时间 : 2020/3/21 20:46
** @作者 : SongZhiBin
******/
package main

import "fmt"

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
func main() {
	//out := GenerateNatural()
	out := make(chan int)
	go func() {
		for i := 2; ; i++ {
			out <- i
			fmt.Printf("111 :%p\n", out)
		}
	}()

	//out := make(chan int)
	//go func(out chan int) {
	//	for i := 2; ; i++ {
	//		out <- i
	//	}
	//}(out)
	for i := 0; i < 100; i++ {
		prime := <-out // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		out = PrimeFilter(out, prime) // 基于新素数构造的过滤器
		fmt.Printf("112 :%p\n", out)
	}
}
