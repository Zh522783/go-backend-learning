package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8081/hello")
	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}
	defer resp.Body.Close() 

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败：", err)
		return
	}

	fmt.Println("HTTP状态码：", resp.Status)
	fmt.Println("服务器返回内容：", string(body))
}
