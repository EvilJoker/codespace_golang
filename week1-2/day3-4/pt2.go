/*
声明一个映射（map），将水果名称映射到其价格。添加一些水果到映射中，然后打印出水果和其价格。

运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import "fmt"

func main() {
	var map_f = map[string]int{"apple": 10, "pear": 20}

	for name, price := range map_f {
		fmt.Println(name, price)
	}

}
