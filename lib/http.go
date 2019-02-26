package lib

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var client = http.DefaultClient

// 执行HTTP请求
// method 请求方法：GET,POST,PUT,DELETE
// url 请求地址，必须是一个合法的HTTP地址
// params 请求参数，GET请求时使用nil代替
// headers 请求消息头，注意POST、PUT请求时，必须指定Content-Type响应，其它按需指定
// 返回：响应状态码、响应字节数、错误
func DoRequest(method, url, params string, headers *map[string]string) (string, int, error) {

	// 构建请求
	req, err := http.NewRequest(method, url, strings.NewReader(params))
	if err != nil {
		return "100", 0, err
	}

	// 添加请求消息头
	if headers != nil {
		for name, value := range *headers {
			req.Header.Set(name, value)
		}
	}

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}

	// 关闭响应
	defer resp.Body.Close()

	// 响应状态码
	status := strings.Split(resp.Status, " ")[0]

	// 读取响应
	if data, err := ioutil.ReadAll(resp.Body); err != nil {
		return status, 0, err
	} else {
		return status, len(data), nil
	}

}
