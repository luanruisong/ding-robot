package ding_robot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

//1.把timestamp+"\n"+密钥当做签名字符串
//2.使用HmacSHA256算法计算签名
//3.然后进行Base64 encode
//4.最后再把签名参数再进行urlEncode
//5.得到最终的签名（需要使用UTF-8字符集）

const (
	baseFmt = "%d\n%s"
)

func (dm *DingMsg) SignData(secret string) *SignData {
	sd := &SignData{}
	if len(secret) > 0 {
		sd.Timestamp = Now()
		signStr := fmt.Sprintf(baseFmt, sd.Timestamp, secret)
		str := HmacSha256(signStr, secret)
		sd.Sign = url.QueryEscape(str)
	}
	return sd
}

func Now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(sha)
}
