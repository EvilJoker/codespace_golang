/*
4. 编写一个程序，接受用户输入的数字，如果输入的数字是偶数，则打印 "这是一个偶数"，否则打印 "这是一个奇数"。

运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	if input%2 == 0 {
		fmt.Print("这是一个偶数")
	} else {
		fmt.Print("这是一个奇数")
	}
}
