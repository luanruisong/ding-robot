package ding_robot

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/luanruisong/greq"
)

type (
	DingClient struct {
		Token  string
		Secret string
	}
)

const BaseDingUrl = "https://oapi.dingtalk.com/robot/send"

func (d *DingClient) getFinalUrl(sign *Sign) string {
	url := fmt.Sprintf("%s?access_token=%s", BaseDingUrl, d.Token)
	if sign != nil {
		url += fmt.Sprintf("&timestamp=%d&sign=%s", sign.Timestamp, sign.Sign)
	}
	return url
}

func (d *DingClient) SendRobotDingMsg(msg *DingMsg) (err error) {
	if len(d.Token) == 0 {
		err = fmt.Errorf("can not find token")
	} else {
		sign := SignDataWithNow(d.Secret)
		req := greq.NewJson(d.getFinalUrl(sign))
		resp := req.Post(msg)
		if resp.Ok {
			var resCode = struct {
				ErrorCode int    `json:"errorcode"`
				ErrMsg    string `json:"errmsg"`
			}{}
			if err = jsoniter.Unmarshal(resp.RawBody, &resCode); err == nil {
				if resCode.ErrorCode == 0 {
					logger.Info("ding msg send success")
				} else {
					logger.Warn("ding msg send failed ---> %s", resCode.ErrMsg)
					err = errors.New(resCode.ErrMsg)
				}
			}
		} else {
			logger.Error("ding msg send failed %v", resp.Err)
			err = resp.Err
		}
	}
	return
}

func (d *DingClient) SendTextMsg(msg string, atMobiles ...string) error {
	dMsg := NewTextMsg(msg)
	if len(atMobiles) > 0 {
		dMsg.At(atMobiles...)
	}
	return d.SendRobotDingMsg(dMsg)
}

func (d *DingClient) SendLinkMsg(title, text, imgSrc, href string) error {
	return d.SendRobotDingMsg(NewLinkMsg(title, text, imgSrc, href))
}

func (d *DingClient) SendMarkdownMsg(title, text string) error {
	return d.SendRobotDingMsg(NewMarkdownMsg(title, text))
}

func NewClient(token, secret string) *DingClient {
	return &DingClient{
		Token:  token,
		Secret: secret,
	}
}
