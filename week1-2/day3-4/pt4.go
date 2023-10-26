/*
4. 编写一个函数，接受一个整数切片作为参数，并返回切片中的最大值和最小值。


运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import "fmt"

func max(arr []int) (int, int) {
	max := arr[0]
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	return max, min
}

func main() {
	arr := []int{1, 23, 6, 4, 4}
	max1, min2 := max(arr)
	fmt.Println(max1, min2)
}
