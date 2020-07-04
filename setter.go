package ding_robot

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/luanruisong/greq"
)

func (d *DingClient) getFinalUrl(msg *SignData) string {
	return fmt.Sprintf("%s?access_token=%s&timestamp=%d&sign=%s", d.Url, d.Token, msg.Timestamp, msg.Sign)
}

func (d *DingClient) SendRobotDingMsg(msg *DingMsg) {
	if len(d.Token) == 0 || len(d.Url) == 0 {
		json, _ := jsoniter.MarshalToString(d)
		logger.Error("can not send ding msg by", json)
		return
	}
	sign := msg.SignData(d.Secret)
	req := greq.NewJson(d.getFinalUrl(sign))
	resp := req.Post(msg)
	if resp.Ok {
		var rescode = struct {
			Errorcode int    `json:"errorcode"`
			Errmsg    string `json:"errmsg"`
		}{}
		_ = jsoniter.Unmarshal(resp.RawBody, &rescode)
		if rescode.Errorcode == 0 {
			logger.Info("ding msg send success")
		} else {
			logger.Warn("ding msg send failed ---> %s", rescode.Errmsg)
		}

	} else {
		logger.Error("ding msg send failed %v", resp.Err)
	}
}

func (DefaultLogger) Info(s string, i ...interface{}) {
	fmt.Println(fmt.Sprintf(s, i...))
}
func (DefaultLogger) Warn(s string, i ...interface{}) {
	fmt.Println(fmt.Sprintf(s, i...))
}
func (DefaultLogger) Error(s string, i ...interface{}) {
	fmt.Println(fmt.Sprintf(s, i...))
}
