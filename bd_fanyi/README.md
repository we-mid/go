# bd_fanyi

> 简单易用的 百度翻译 SDK

```ini
# .env
BAIDUFANYI_APP_ID=xxxxxxxx
BAIDUFANYI_APP_SECRET=xxxxxxxx
```

```go
import "gitee.com/we-mid/go/bd_fanyi"

func init() {
	bd_fanyi.InitFromEnv()
}

func handleFanyi(w http.ResponseWriter, r *http.Request) (any, error) {
	defer r.Body.Close()
	// ...
	var req bd_fanyi.FanyiReq
	if err := json.NewDecoder(r.Body).Decode(&req); err!=nil {
		return nil, fmt.Errorf("json.Decode: %w", err)
	}
	bd_fanyi.CompleteReq(&req)
	// list: []*bd_fanyi.TranslateResult
	list, err := client.Translate(&req)
	if err != nil {
		return nil, fmt.Errorf("client.Translate: %w", err)
	}
	return list, nil
}
```

```go
// bd_fanyi.FanyiReq:
type FanyiReq struct {
	Text string      `json:"text"`
	From baidufanyi.Language `json:"from"`
	To   baidufanyi.Language `json:"to"`
}
// bd_fanyi.Language = baidufanyi.Language:
// https://github.com/chyroc/baidufanyi/blob/master/language.go
LanguageEn  Language = "en"  // 英
LanguageCht Language = "cht" // 中文(繁体)
LanguageWyw Language = "wyw" // 中文(文言文)
LanguageYue Language = "yue" // 中文(粤语)
LanguageZh  Language = "zh"  // 中文(简体)

// bd_fanyi.TranslateResult = baidufanyi.TranslateResult:
// https://github.com/chyroc/baidufanyi/blob/master/translate.go
type TranslateResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}
```

### 参考链接

- 通用翻译API接入文档：https://fanyi-api.baidu.com/doc/21
- 通用翻译API版本权益调整通知（计费标准）：https://fanyi-api.baidu.com/doc/8
- 本SDK基于/二次封装了baidufanyi：https://github.com/chyroc/baidufanyi
