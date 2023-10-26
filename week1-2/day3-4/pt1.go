/*
创建一个整数切片（slice），将一些整数值添加到切片中，并计算切片的总和。

运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, 3, 4}
	// append 必须这么干
	arr = append(arr, 5, 6, 7)

	var sum int = 0
	for _, v := range arr {
		sum += v
	}

	fmt.Println(sum)

}
