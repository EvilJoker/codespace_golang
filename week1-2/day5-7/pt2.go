/*
修改上述程序，使用channel来协调两个goroutines的输出。一个goroutine生成奇数并将它们发送到一个奇数通道，另一个生成偶数并将它们发送到一个偶数通道。主goroutine从这两个通道接收数字并打印。

运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		for i := 1; i < 10; i += 2 {
			chan1 <- i
		}
		close(chan1)
	}()
	go func() {
		for i := 0; i < 10; i += 2 {
			chan2 <- i
		}
		close(chan2)
	}()

	value1, ok1 := <-chan1
	value2, ok2 := <-chan2

	for ok1 || ok2 {
		if ok1 {
			fmt.Println(value1)
			value1, ok1 = <-chan1
		}
		if ok2 {
			fmt.Println(value2)
			value2, ok2 = <-chan2
		}
	}

}
