package lib

import "testing"

// 在 lib 目录下执行测试：go test -v
// === RUN   Test_DoRequest
// --- PASS: Test_DoRequest (3.80s)
//     http_test.go:13: Test passed, response bytes is 260
// PASS
// ok      github.com/zlikun/httpbench/lib 4.694s
func Test_DoRequest(t *testing.T) {
	status, size, err := DoRequest("GET", "http://httpbin.org/get?v=1.0", "",
		&map[string]string{"User-Agent": "go version go1.11.2 windows/amd64"})

	if err != nil {
		t.Error(err)
	} else {
		if status == "200" {
			t.Logf("Test passed, response bytes is %v", size)
		} else {
			t.Errorf("Response status is %v", status)
		}
	}
}
