package main

import (
	"fmt"
	"github.com/zlikun/httpbench/lib"
)

func main() {

	status, size, err := lib.DoRequest("GET", "http://httpbin.org/get?keyword=john&version=1.0", "", nil)
	fmt.Println(status, size, err)

}
