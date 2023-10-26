/*
5. 修改上述生产者-消费者程序，添加多个生产者和多个消费者。确保多个goroutines可以安全地访问共享的通道。

运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	chan1 := make(chan rune, 10)

	wg := sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < 5; i++ {
		// 生产者
		go func(i int) {
			for {
				chan1 <- 'a' + rune(i)
			}
		}(i)
	}

	for i := 0; i < 5; i++ {
		// 消费者
		go func(i int) {
			for {
				time.Sleep(1 * time.Second)
				t := <-chan1
				fmt.Println(i, ":", string(t))
			}
		}(i)
	}
	wg.Wait()

}
