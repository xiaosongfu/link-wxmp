package main

import (
	"log"
	"net/http"
	"time"
)

const token = "link1205fun"
const port = ":80" // 1205 or 80

func main() {
	// 配置服务
	server := &http.Server{
		Addr:           port,
		Handler:        &SimpleServeMux{},
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 启动服务
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
