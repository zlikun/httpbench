package main

import (
	"fmt"
	"github.com/hashicorp/go-uuid"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

var client = http.DefaultClient

// GET请求，返回响应状态码、响应消息字节数、错误信息
func fetch(url string) (status int, length int, err error) {

	// 构建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, 0, err
	}

	// 添加请求消息头
	req.Header.Set("User-Agent", "go version go1.11.2 windows/amd64")

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	} else {
		defer resp.Body.Close()
	}

	// 读取响应
	if data, err := ioutil.ReadAll(resp.Body); err != nil {
		return resp.StatusCode, 0, err
	} else {
		//write(data) // 将响应写入文件，检查请求是否正常
		return resp.StatusCode, len(data), nil
	}
}

func write(data []byte) {
	// 使用UUID作为文件名称
	name, _ := uuid.GenerateUUID()
	// C:\Users\zhanglk\AppData\Local\Temp
	file := path.Join(os.TempDir(), "zlikun-"+name+".log")
	// 文件存放于系统临时目录下的httpbench子目录下
	if err := ioutil.WriteFile(file, data, 0644); err != nil {
		fmt.Println("写入文件出错!", err)
	}
}

type Response struct {
	status int
	length int
	err    error
}

// 测试函数，使用协程并发访问
func benchmark(url string, ch chan<- *Response) {
	status, length, err := fetch(url)
	ch <- &Response{status, length, err}
}

func statistic(ch <-chan *Response) {
	var success int // 成功数
	var error int   // 失败数
	var size int64  // 总字节数
	var begin = time.Now()
	for i := 0; i < requests; i++ {
		resp := <-ch
		if resp.err != nil || resp.status != 200 {
			error++
		} else {
			success++
			size += int64(resp.length)
		}
	}
	during := time.Now().Sub(begin)

	fmt.Printf("并发[%v]访问[%v]，共%v次，成功%v次，失败%v次，总耗时：%v，总字节数：%v，统计：%.2f req/sec\n",
		users, url, requests, success, error, during, size, float64(requests)/during.Seconds())
}

var url = "https://zlikun.com"
var users = 50       // 并发数
var requests = 10000 // 请求数

// 并发GET请求
func main() {

	ch := make(chan *Response, users)

	// 并发访问，并发数：20
	for i := 0; i < requests; i++ {
		go benchmark(url, ch)
	}

	// 统计计数
	statistic(ch)

}
