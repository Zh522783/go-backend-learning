package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/hello" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintln(w, "Hello, World!")
	})

	fmt.Println("服务启动，监听端口 :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("服务启动失败：", err)
	}
}
