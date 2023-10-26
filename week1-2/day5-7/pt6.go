/*
创建一个简单的Web服务器，它能够并发处理多个HTTP请求。每个请求都会触发一个独立的goroutine来处理。您可以使用Go的`net/http`包来构建这个服务器。


运行 ：
go run xxx.go

备注 ：
单文件执行，忽略 main 告警
*/

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	goroutineID := rand.Intn(100)
	fmt.Printf("Goroutine ID: %d\n", goroutineID)

}

func main() {

	// 并发访问
	go func() {
		for {
			time.Sleep(1 * time.Second)
			http.Get("http://127.0.0.1:8080/")
		}
	}()
	http.HandleFunc("/", handler)
	// 对于每个请求 会自动创建一个协程，处理 请求
	http.ListenAndServe(":8080", nil)

}
