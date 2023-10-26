/*
5. 编写一个程序，模拟银行账户的交易。您可以定义一个账户结构，包括账户余额和一些交易方法，如存款和取款。

运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import "fmt"

type Account struct {
	money float32
}

func saveMoney(acc *Account, money float32) {
	acc.money += money
	fmt.Println("save money success, sum", acc.money)
}

func takeMoney(acc *Account, money float32) {
	acc.money -= money
	fmt.Println("take money success, sum", acc.money)
}

func main() {
	var account Account = Account{0.0}

	saveMoney(&account, 100.0)
	takeMoney(&account, 19.0)
}
