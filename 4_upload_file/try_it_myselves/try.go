package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/upload", handleUpload)

	http.ListenAndServe("127.0.0.1:9090", nil)
}

func handleUpload(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		files, _ := template.ParseFiles("4_upload_file/upload.gtpl")
		files.Execute(writer, nil)
	}

	if request.Method == http.MethodPost {
		file, handler, e := request.FormFile("uploadfile")
		if e != nil {
			fmt.Fprintln(writer, "上传失败")
			return
		}
		defer file.Close()
		fmt.Fprintf(writer, "%v", handler.Header)
		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
