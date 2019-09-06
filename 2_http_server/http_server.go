package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 匹配路由，并传一个处理函数
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/user/login", handleLogin)

	// 启动服务并监听地址，传nil
	err := http.ListenAndServe("127.0.0.1:8880", nil)
	if err != nil {
		fmt.Println("failed: ", err)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	// 接收参数，默认不会接收
	_ = r.ParseForm()
	fmt.Println("所有的参数", r.Form)
	fmt.Println("Host", r.URL.Host)
	fmt.Println("Path", r.URL.Path)
	fmt.Println("welcome to goland world!")
	// 这个是向网页输出的内容
	_, _ = fmt.Fprintf(w, "hello")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request login")
	// 这个是向网页输出的内容
	_, _ = fmt.Fprintf(w, "login")
}
