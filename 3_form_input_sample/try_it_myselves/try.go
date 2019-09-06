package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", hangLogin)

	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		fmt.Println("error occurred: ", err)
	}
}

func hangLogin(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		t, _ := template.ParseFiles("3_form_input_sample/try_it_myselves/login.hello")

		// 生成md5加密的时间戳token
		crutime := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		_ = t.Execute(writer, token)
	}
	// 使用FormValue
	if request.Method == http.MethodPost {
		name := request.FormValue("username")
		password := request.FormValue("password")
		if len(name) == 0 {
			fmt.Fprintln(writer, "用户名不能为空")
		} else if m, _ := regexp.MatchString("^[0-9]+$", name); !m {
			fmt.Fprintln(writer, "用户名必须为纯数字")
		} else {
			_, _ = fmt.Fprintln(writer, "username_: "+name)
		}

		if len(password) == 0 {
			fmt.Fprintln(writer, "密码不能为空")
		} else {
			_, _ = fmt.Fprintln(writer, "password_: "+password)
		}

	}
}

func handleHome(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "welcome to homepage")
}
