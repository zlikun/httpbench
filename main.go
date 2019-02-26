package main

import (
	"flag"
	"fmt"
	"github.com/zlikun/httpbench/lib"
	"log"
	"os"
	"time"
)

var usage = `HTTP Benchmark Tool

This is an HTTP benchmark tool.

Usage: httpbench [options] [http[s]://]hostname[:port]/path
`

var (
	help        bool   // 是否打印帮助信息
	concurrency int    // 并发数
	requests    int    // 请求数
	url         string // 请求地址
)

func init() {

	flag.BoolVar(&help, "h", false, "This help")
	flag.IntVar(&concurrency, "c", 0, "Number of multiple requests to make at a time")
	flag.IntVar(&requests, "n", 0, "Number of requests to perform")
	flag.StringVar(&url, "url", "", "Request url")

	flag.Usage = func() {
		fmt.Println(usage)
		flag.PrintDefaults()
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Data struct {
	status string
	size   int
	err    error
}

func fetch(ch chan *Data) {
	status, size, err := lib.DoRequest("GET", url, "", nil)
	ch <- &Data{
		status: status,
		size:   size,
		err:    err,
	}
}

func monitor(ch chan *Data) {
	var success, failure int
	var begin = time.Now()
	for i := 0; i < requests; i++ {
		data := <-ch
		if data.err != nil {
			failure++
		} else {
			success++
		}
	}

	during := time.Now().Sub(begin)

	fmt.Printf("并发数：%v，请求数：%v，成功次数：%v，失败次数：%v，总耗时：%v，平均：%.2f req/sec\n",
		concurrency, requests, success, failure, during, float64(requests)/during.Seconds())

	os.Exit(0)
}

func main() {

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	// 使用管道缓冲区大小控制协程并发
	var ch = make(chan *Data, concurrency)

	// 暂时只支持GET请求，且不能定义Header
	// 循环调用协程函数，一共请求 requests 次
	for i := 0; i < requests; i++ {
		go fetch(ch)
	}

	// 处理响应（统计、计数）
	// 并发数：20，请求数：1000，成功次数：768，失败次数：232，总耗时：38.4561202s，平均：26.003663260861142 request/second
	monitor(ch)

	fmt.Println("Exit ...")
}
