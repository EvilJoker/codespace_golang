/*
1. 创建一个简单的Go程序，其中启动两个goroutines。一个goroutine打印奇数，另一个打印偶数，打印从1到10的数字。确保它们交替打印。
交替打印用管道，无缓冲打印是阻塞
运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import (
	"fmt"
	"sync"
)

// 互相阻塞，只有另外一个管道有值时，才会进行

func print1(wg *sync.WaitGroup, c1 chan int, c2 chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			if i != 9 {
				<-c2 // 读取 丢弃
				fmt.Println(i)
				c1 <- 1
			} else {
				// 最后一个执行的关闭管道，否则 c2 无法消费
				<-c2 // 读取 丢弃
				fmt.Println(i)
				close(c1)
			}

		}
	}
	// <-c2 // 读取 丢弃
	// c1 <- 1

}

func print2(wg *sync.WaitGroup, c1 chan int, c2 chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {

		if i%2 == 0 {

			if i != 10 {
				<-c1 // 读取 丢弃
				fmt.Println(i)
				c2 <- 1
			} else {
				// 最后一个执行的关闭管道，否则 c2 无法消费
				<-c1 // 读取 丢弃
				fmt.Println(i)
				close(c2)
			}
		}
	}
}

func main() {
	// 主程序不会等协程结束，需要加锁
	var wg sync.WaitGroup
	wg.Add(2)

	chan1 := make(chan int)
	chan2 := make(chan int)

	go print1(&wg, chan1, chan2)
	go print2(&wg, chan1, chan2)
	// go print1(&wg, chan1, chan2)
	chan2 <- 0

	wg.Wait()

}
