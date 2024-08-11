// Origin: https://github.com/we-mid/bec-services/tree/main/fanyi
package bd_fanyi

import (
	"os"

	bf "github.com/chyroc/baidufanyi"
)

// 通用翻译API接入文档
// https://fanyi-api.baidu.com/doc/21

// 通用翻译API版本权益调整通知
// https://fanyi-api.baidu.com/doc/8
// 标准版 QPS=1 持28个语种互译 单次最长请求1000字符 免费调用量5万字符/月
// 高级版 QPS=10 支持 28个语种互译 单次最长请求6000字符 免费调用量100万字符/月
// 尊享版 QPS=100 支持200+语种互译 单次最长请求6000字符 免费调用量200万字符/月
// 2. 若每月调用量超过免费调用量限制，将按照49元/百万字符进行计费；
// 3. 字符数以翻译的源语言字符长度为标准计算。空格、html标签等均计入字符。一个汉字，英文字母，标点符号等，均按照一个字符计费。
var client *bf.Fanyi

func InitFromEnv() {
	appid := os.Getenv("BAIDUFANYI_APP_ID")
	secret := os.Getenv("BAIDUFANYI_APP_SECRET")
	client = bf.New(bf.WithCredential(appid, secret))
}

// alias
type TranslateResult = bf.TranslateResult
type Language = bf.Language

type FanyiReq struct {
	Text string      `json:"text"`
	From bf.Language `json:"from"`
	To   bf.Language `json:"to"`
}

func Translate(req *FanyiReq) ([]*bf.TranslateResult, error) {
	CompleteReq(req)
	return client.Translate(req.Text, req.From, req.To)
}

func CompleteReq(req *FanyiReq) {
	// TODO: sdk support language=`auto`
	// TODO: auto detect language from text?
	if req.From == "" && req.To == "" {
		req.From = bf.LanguageEn
		req.To = bf.LanguageZh
	} else if req.From == "" {
		if req.To == bf.LanguageEn {
			req.From = bf.LanguageZh
		} else {
			req.From = bf.LanguageEn
		}
	} else if req.To == "" {
		if req.From == bf.LanguageEn {
			req.To = bf.LanguageZh
		} else {
			req.To = bf.LanguageEn
		}
	}
}
