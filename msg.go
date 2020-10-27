package ding_robot

import "fmt"

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

func (dm *DingMsg) initAt() error {
	if dm.MsgType != MsgTypeText {
		return fmt.Errorf("type:%s can not add at", dm.MsgType)
	}
	if dm.At_ == nil {
		dm.At_ = &DingMsgAt{
			AtMobiles: nil,
			IsAtAll:   false,
		}
	}
	return nil
}

func (dm *DingMsg) At(mobiles ...string) (err error) {
	if err = dm.initAt(); err == nil {
		dm.At_.AtMobiles = append(dm.At_.AtMobiles, mobiles...)
	}
	return
}

func (dm *DingMsg) AtAll() (err error) {
	if err = dm.initAt(); err == nil {
		dm.At_.IsAtAll = true
	}
	return
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
