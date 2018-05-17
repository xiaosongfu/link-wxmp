package main

import (
	"github.com/xiaosongfu/link-wxmp/wechat"
	"net/http"
	"regexp"
)

func init() {
	mux = append(mux, WebController{HandlerFunc: logger(wxApiGet), Method: http.MethodGet, Pattern: "/wxapi"})
	mux = append(mux, WebController{HandlerFunc: logger(wxApiPost), Method: http.MethodPost, Pattern: "/wxapi"})
}

// 控制器
type WebController struct {
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	Method      string
	Pattern     string
}

var mux []WebController

// 路由
type SimpleServeMux struct{}

// 实现 ServeHTTP 方法
func (simpleServeMux *SimpleServeMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, webController := range mux {
		if match, _ := regexp.MatchString(webController.Pattern, request.URL.Path); match {
			if request.Method == webController.Method {
				webController.HandlerFunc(writer, request)
				return // case 1: 匹配上了
			}
		}
	}
	// case 2: 没有匹配上
	writer.Write([]byte(""))
}

// 处理 Get 请求的方法
func wxApiGet(writer http.ResponseWriter, request *http.Request) {
	auth, err := wechat.Authentication(request, token)
	if err == nil && len(auth.EchoStr) > 0 {
		writer.Write([]byte(auth.EchoStr))
		return
	}
	writer.WriteHeader(403)
}

// 处理 Post 请求的方法
func wxApiPost(writer http.ResponseWriter, request *http.Request) {
	// 判断签名
	_, err := wechat.Authentication(request, token)
	if err != nil {
		writer.Write([]byte(err.Error()))
		writer.WriteHeader(403)
		return
	}
	// 处理请求
	wechat.ProcessReceiveMessage(writer, request)
}
