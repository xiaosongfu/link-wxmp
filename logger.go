package main

import (
	"log"
	"net/http"
	"time"
)

func logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// 开始时间
		start := time.Now()

		h.ServeHTTP(writer, request)

		// 打印日志
		log.Printf(
			"%s\t%s\t%s",
			request.Method,
			request.RequestURI,
			time.Since(start),
		)
	})
}
