package wechat

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"net/http"
	"sort"
)

type Auth struct {
	Signature    string `json:"signature"`
	Timestamp    string `json:"timestamp"`
	Nonce        string `json:"nonce"`
	EncryptType  string `json:"encrypt_type"`
	MsgSignature string `json:"msg_signature"`
	EchoStr      string `json:"echostr"`
}

// 实现微信认证
func Authentication(request *http.Request, token string) (auth Auth, err error) {
	// 获取微信服务器传过来的测试
	auth.Nonce = request.URL.Query().Get("nonce")
	auth.EchoStr = request.URL.Query().Get("echostr")
	auth.Signature = request.URL.Query().Get("signature")
	auth.Timestamp = request.URL.Query().Get("timestamp")
	auth.EncryptType = request.URL.Query().Get("encrypt_type")
	auth.MsgSignature = request.URL.Query().Get("msg_signature")

	// 加密签名
	signature := auth.generateWeChatSignature(token)
	if auth.Signature == signature {
		err = nil
		return
	}
	err = errors.New("Invalid Signature")
	return
}

// 生成加密签名
func (auth *Auth) generateWeChatSignature(token string) (signature string) {
	// 排序
	strSlice := sort.StringSlice{token, auth.Timestamp, auth.Nonce}
	sort.Strings(strSlice)

	// 把 string slice 转为 string
	str := ""
	for _, s := range strSlice {
		str += s
	}

	// sha1 加密
	sh := sha1.New()
	sh.Write([]byte(str))
	return fmt.Sprintf("%x", sh.Sum(nil))
}
