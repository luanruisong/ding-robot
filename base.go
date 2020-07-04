package ding_robot

const (
	MSG_TYPE = "text"
)

var logger Logger

func init() {
	logger = &DefaultLogger{}
}

func SetLogger(l Logger) {
	logger = l
}

func NewClient(token, secret string) *DingClient {
	return &DingClient{
		Url:    "https://oapi.dingtalk.com/robot/send",
		Token:  token,
		Secret: secret,
	}
}

func NewTextMsg(msg string) *DingMsg {
	return &DingMsg{
		Msgtype: MSG_TYPE,
		Text:    &DingMsgContent{msg},
	}
}
