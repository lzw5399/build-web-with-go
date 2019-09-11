package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 反序列化，结构体接收
	// testUnmarshal_struct()

	// 反序列化，interface接收，源生方法
	testUnmarshal_interface_standrad()
}

func testUnmarshal_struct() {
	var s ToolStruct
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	_ = json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

type ToolStruct struct {
	Servers []struct {
		ServerIP   string `json:"serverIP"`
		ServerName string `json:"serverName"`
	} `json:"servers"`
}

func testUnmarshal_interface_standrad() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	_ = json.Unmarshal(b, &f)

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

