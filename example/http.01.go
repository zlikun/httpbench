package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 测试GET请求
func main() {
	var url = "https://zlikun.com"
	var client = http.DefaultClient

	// 构建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("构建请求对象出错!", err)
	}

	// 添加请求消息头
	req.Header.Set("User-Agent", "go version go1.11.2 windows/amd64")

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("执行请求出错!", err)
	} else {
		defer resp.Body.Close()
	}

	// 响应状态码
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)

	// 读取响应
	if data, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Fatal("读取响应消息出错!", err)
	} else {
		fmt.Println(string(data))
	}
}
