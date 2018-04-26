package wechat

import "regexp"

// 封装 CDATA 类型 xml 数据
func value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

// 判断 url 是否合法
func urlIsCorrect(url string) (match bool) {
	matchHttp, _ := regexp.MatchString("^http://", url)
	matchHttps, _ := regexp.MatchString("^https://", url)
	return matchHttp || matchHttps
}
