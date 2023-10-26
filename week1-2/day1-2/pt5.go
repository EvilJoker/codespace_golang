/*
5. 使用 `for` 循环打印出 1 到 10 的整数。

运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}
