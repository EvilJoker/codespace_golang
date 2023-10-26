/*
创建一个自定义的结构体（struct），表示一位学生的信息，包括姓名、年龄和分数。使用结构体创建多个学生的对象，并打印他们的信息。

运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import "fmt"

type Student struct {
	name  string
	age   int
	score float64
}

func main() {
	var stu1 Student = Student{"zhangsan", 18, 99.5}
	var stu2 Student = Student{"zhangsan2", 18, 99.5}
	var stu3 Student = Student{"zhangsan3", 18, 99.5}

	var stu_arr []Student
	stu_arr = append(stu_arr, stu1)
	stu_arr = append(stu_arr, stu2)
	stu_arr = append(stu_arr, stu3)

	for _, stu := range stu_arr {
		fmt.Println(stu.name, stu.age, stu.score)
	}

}
