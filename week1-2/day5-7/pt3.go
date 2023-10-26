/*
3. 创建一个带缓冲的channel，向其中发送一些值。然后启动一个goroutine来从该channel接收值，但是在接收前等待一段时间。观察带缓冲channel如何处理等待的情况。

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
	chan1 := make(chan int, 2)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for c := range chan1 {
			time.Sleep(1 * time.Second)
			fmt.Println("output", c)
		}

	}()

	for i := 0; i < 10; i++ {
		fmt.Println("input", i)
		chan1 <- i
	}
	close(chan1)
	wg.Wait()

}
