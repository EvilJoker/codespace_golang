/*
创建一个字符串切片，包含几个水果的名称，然后使用 `for` 循环打印出每个水果的名称。

运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import "fmt"

func main() {
	// 和数组的区别是没有固定大小
	var friuts = []string{"apple", "banana", "orange"}

	// 和 map 的区别是 key 是序号
	for _, item := range friuts {
		fmt.Println(item)
	}
}
