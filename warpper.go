package ding_robot

type (

	//"at": {
	//		"atMobiles": [
	//		"156xxxx8827",
	//		"189xxxx8325"
	//	],
	//	"isAtAll": false
	//}

	DingMsgContent struct {
		Content string `json:"content,omitempty"`
	}
	DingMsgAt struct {
		AtMobiles []string `json:"atMobiles,omitempty"`
		IsAtAll   bool     `json:"isAtAll,omitempty"`
	}
	DingMsg struct {
		Msgtype string          `json:"msgtype,omitempty"`
		Text    *DingMsgContent `json:"text,omitempty"`
		At      *DingMsgAt      `json:"at,omitempty"`
	}
	SignData struct {
		Timestamp int64  `json:"timestamp,omitempty"`
		Sign      string `json:"sign,omitempty"`
	}

	DingClient struct {
		Url    string
		Token  string
		Secret string
	}

	Logger interface {
		Info(s string, i ...interface{})
		Warn(s string, i ...interface{})
		Error(s string, i ...interface{})
	}

	DefaultLogger struct {
	}
)
