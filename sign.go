package ding_robot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

type (
	Sign struct {
		Timestamp int64  `json:"timestamp,omitempty"`
		Sign      string `json:"sign,omitempty"`
	}
)

const (
	baseFmt = "%d\n%s"
)

func SignDataWithNow(secret string) *Sign {
	return SignData(Now(), secret)
}
func SignData(ts int64, secret string) *Sign {

	//1.把timestamp+"\n"+密钥当做签名字符串
	//2.使用HmacSHA256算法计算签名
	//3.然后进行Base64 encode
	//4.最后再把签名参数再进行urlEncode
	//5.得到最终的签名（需要使用UTF-8字符集）
	if len(secret) > 0 && ts > 0 {
		sd := &Sign{}
		sd.Timestamp = Now()
		signStr := fmt.Sprintf(baseFmt, sd.Timestamp, secret)
		str := HmacSha256(signStr, secret)
		sd.Sign = url.QueryEscape(str)
		return sd
	}
	return nil
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
