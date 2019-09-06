package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", handleLogin)

	m := &MyMux{}
	_ = http.ListenAndServe("127.0.0.1:8880", m)
}

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		_, _ = fmt.Fprintf(w, "hello from custom handler")
		return
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "login")
}
