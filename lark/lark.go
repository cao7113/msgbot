package lark

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type BotClient struct {
	CustomBot
	c resty.Client
}

const (
	MsgText = "text"
	MsgPost = "post"
)

type PostMsg struct {
}

type CustomBot struct {
	Webhook   string
	SigSecret string
}

func (b *CustomBot) sign() (string, error) {
	return GenSign(b.SigSecret, time.Now().Unix())
}

// GenSign https://www.feishu.cn/hc/zh-CN/articles/360024984973#magicdomid-3_91
func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
