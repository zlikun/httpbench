# HTTP Benchmark Tool

![license](https://img.shields.io/github/license/alibaba/dubbo.svg)

> 类似 `ab` HTTP API 压测工具
```
# 帮助文档
$ httpbench -h
HTTP Benchmark Tool

This is an HTTP benchmark tool.

Usage: httpbench [options] [http[s]://]hostname[:port]/path

  -h    This help
  -c int
        Number of multiple requests to make at a time
  -n int
        Number of requests to perform
  -url string
        Request url

# 演示用例
$ httpbench -c 20 -n 1000 -url https://www.baidu.com
  并发数：20，请求数：1000，成功次数：1000，失败次数：0，总耗时：4.2207134s，平均：236.93 req/sec
```

