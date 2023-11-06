package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

// 测试 net/http
func main() {

	//新建一个路径映射处理器{"/foo", 自定义函数}
	http.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
	})

	//使用http开启服务，并监听127.0.0.1:8080，监听到后，对应的处理器是nil
	log.Fatal(http.ListenAndServe(":8080", nil))
}
