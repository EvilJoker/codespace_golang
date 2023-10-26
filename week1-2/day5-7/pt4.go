/*
使用goroutines模拟一个简单的生产者-消费者模型。创建一个字符串通道，一个goroutine不断向通道发送随机生成的字符串，另一个goroutine从通道接收并打印这些字符串。

运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	chan1 := make(chan rune, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			randomInt := rand.Intn(26)
			randomLetter := 'a' + rune(randomInt)
			chan1 <- randomLetter
		}
		close(chan1)
	}()
	go func() {
		defer wg.Done()

		for randomLetter := range chan1 {
			fmt.Printf("%c\n", rune(randomLetter))
		}

	}()

	wg.Wait()

}
