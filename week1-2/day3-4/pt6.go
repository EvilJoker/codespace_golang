/*
创建一个函数，接受一个整数值 n 作为参数，然后打印出一个 n x n 的乘法表。


运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import "fmt"

func printMultiplicationTable(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d*%d=%d\t", i+1, j+1, (i+1)*(j+1))
		}
	}
}

func main() {
	printMultiplicationTable(5)

}
