package wechat

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/clbanning/mxj"
	"github.com/xiaosongfu/link-wxmp/link"
	"io/ioutil"
	"net/http"
	"strings"
)

type PostRequest struct {
	Message map[string]interface{}
}

// 处理请求
func ProcessReceiveMessage(writer http.ResponseWriter, request *http.Request) {
	// PostRequest 实例
	var postRequest PostRequest
	// 解析收到的消息
	err := postRequest.initReceiveMessage(request)
	if err != nil {
		writer.WriteHeader(403)
		return
	}
	// 解析收到的消息的类型
	msgType, ok := postRequest.Message["MsgType"].(string)
	if !ok {
		writer.WriteHeader(403)
		return
	}

	// 判断消息的类型,并执行对应的业务逻辑
	switch msgType {
	case "text":
		postRequest.handleReceiveTextMessage(writer)
	}
	return
}

// 初始化收到的数据
func (postRequest *PostRequest) initReceiveMessage(request *http.Request) error {
	// 读取 post 请求的 body
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.New("invalid request")
	}
	// 解析 xml 数据
	xmlMap, err := mxj.NewMapXml(body)
	if _, ok := xmlMap["xml"]; !ok {
		return errors.New("invalid message")
	}
	msg, ok := xmlMap["xml"].(map[string]interface{})
	if !ok {
		return errors.New("invalid field `xml` type")
	}
	postRequest.Message = msg
	return nil
}

//-------------------- *** --------------------
//-------------------- *** --------------------
// 处理消息 --> 文本类型
func (postRequest *PostRequest) handleReceiveTextMessage(writer http.ResponseWriter) {
	inTextMsg, ok := postRequest.Message["Content"].(string)
	if !ok {
		writer.WriteHeader(403)
		return
	}

	// 回复的内容
	var replyTextContent string

	// 处理 url 并保存
	if urlIsCorrect(inTextMsg) {
		var r http.Request
		r.ParseForm()
		r.Form.Add("url", inTextMsg)
		bodyStr := strings.TrimSpace(r.Form.Encode())
		resp, err := http.Post("http://120.77.47.141:1205/api/v1/addLink", "application/x-www-form-urlencoded", strings.NewReader(bodyStr))
		if err != nil {
			replyTextContent = link.RequestServerError
		} else {
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				replyTextContent = link.RequestServerError
			} else {
				var linkResp = link.Response{}
				err = json.Unmarshal(bytes, &linkResp)
				if err != nil {
					replyTextContent = link.RequestServerError
				} else {
					if linkResp.Code == link.RequestSuccessCode {
						replyTextContent = link.RequestSuccess
					} else {
						replyTextContent = linkResp.Message
					}
				}
			}
		}
	} else {
		replyTextContent = link.RequestUrlError
	}

	// 回复微信服务器的消息结构
	var replayTextMsg TextMessage
	replayTextMsg.InitBaseMessage(postRequest, MsgTypeText)
	replayTextMsg.Content = value2CDATA(replyTextContent)

	// 封装成 xml 格式
	replyXml, err := xml.Marshal(replayTextMsg)
	if err != nil {
		writer.WriteHeader(403)
		return
	}
	// 发送回复
	writer.Header().Set("Content-Type", "text/xml")
	writer.Write(replyXml)
}
