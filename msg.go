package ding_robot

const (
	MsgTypeText     = "text"
	MsgTypeLink     = "link"
	MsgTypeMarkdown = "markdown"
)

type (
	TextMsg struct {
		Content string `json:"content,omitempty"`
	}
	LinkMsg struct {
		Text       string `json:"text"`
		Title      string `json:"title"`
		PicUrl     string `json:"picUrl"`
		MessageUrl string `json:"messageUrl"`
	}
	MarkdownMsg struct {
		Text  string `json:"text"`
		Title string `json:"title"`
	}
	DingMsgAt struct {
		AtMobiles []string `json:"atMobiles,omitempty"`
		IsAtAll   bool     `json:"isAtAll,omitempty"`
	}
	DingMsg struct {
		MsgType  string       `json:"msgtype"`
		Text     *TextMsg     `json:"text,omitempty"`
		Link     *LinkMsg     `json:"link,omitempty"`
		Markdown *MarkdownMsg `json:"markdown,omitempty"`
		At_      *DingMsgAt   `json:"at,omitempty"`
	}
)

func (dm *DingMsg) initAt() {
	if dm.At_ == nil {
		dm.At_ = &DingMsgAt{
			AtMobiles: nil,
			IsAtAll:   false,
		}
	}
}

func (dm *DingMsg) At(mobiles ...string) {
	dm.initAt()
	dm.At_.AtMobiles = append(dm.At_.AtMobiles, mobiles...)
}

func (dm *DingMsg) AtAll() {
	dm.initAt()
	dm.At_.IsAtAll = true
}

func NewTextMsg(msg string) *DingMsg {
	return &DingMsg{
		MsgType: MsgTypeText,
		Text:    &TextMsg{msg},
	}
}

func NewLinkMsg(title, text, imgSrc, href string) *DingMsg {
	return &DingMsg{
		MsgType: MsgTypeLink,
		Link: &LinkMsg{
			Text:       text,
			Title:      title,
			PicUrl:     imgSrc,
			MessageUrl: href,
		},
	}
}

func NewMarkdownMsg(title, text string) *DingMsg {
	return &DingMsg{
		MsgType: MsgTypeMarkdown,
		Markdown: &MarkdownMsg{
			Text:  text,
			Title: title,
		},
	}
}
