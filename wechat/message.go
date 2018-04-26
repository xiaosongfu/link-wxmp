package wechat

import (
	"encoding/xml"
	"strconv"
	"time"
)

const MsgTypeText = "text"

type CDATAText struct {
	Text string `xml:",innerxml"`
}

type BaseMessage struct {
	FromUserName CDATAText
	ToUserName   CDATAText
	CreateTime   CDATAText
	MsgType      CDATAText
}

func (baseMsg *BaseMessage) InitBaseMessage(postRequest *PostRequest, msgType string) {
	baseMsg.FromUserName = value2CDATA(postRequest.Message["ToUserName"].(string))
	baseMsg.ToUserName = value2CDATA(postRequest.Message["FromUserName"].(string))
	baseMsg.CreateTime = value2CDATA(strconv.FormatInt(time.Now().Unix(), 10))
	baseMsg.MsgType = value2CDATA(msgType)
}

//-------------------- *** --------------------
//-------------------- *** --------------------
// 文本类型消息
type TextMessage struct {
	BaseMessage
	Content CDATAText
	XMLName xml.Name `xml:"xml"`
}
