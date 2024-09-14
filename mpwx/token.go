package mpwx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// fix: getAccessToken老接口
// res={ errcode: 40001, errmsg: 'invalid credential, access_token is invalid or not latest rid: xxx' }
// 接口调用凭证 /获取稳定版接口调用凭据
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getStableAccessToken.html
const tokenUrl = "https://api.weixin.qq.com/cgi-bin/stable_token"

type tokenReq struct {
	GrantType    string `json:"grant_type"`
	Appid        string `json:"appid"`
	Secret       string `json:"secret"`
	ForceRefresh bool   `json:"force_refresh"`
}
type tokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type tokenState struct {
	accessToken string
	expiresAt   time.Time
}

var stateMap = make(map[string]tokenState)

func (c *MpwxClient) accessToken(ak *string) error {
	if c.appid == "" || c.secret == "" { // assert
		return fmt.Errorf("appid and secret are both required")
	}
	if cache, ok := stateMap[c.appid]; ok {
		if cache.accessToken != "" && cache.expiresAt.After(time.Now()) {
			*ak = cache.accessToken // cache hit
			return nil
		}
	}
	startTime := time.Now()

	// 设置请求体
	var reqBody bytes.Buffer
	req := tokenReq{
		GrantType: "client_credential",
		Appid:     c.appid,
		Secret:    c.secret,
		// access_token无效或不是最新的，可以通过getStableAccessToken？
		// https://developers.weixin.qq.com/community/develop/doc/0004208ab608983e5610dfa3c66400?jumpto=comment
		ForceRefresh: false,
	}
	if err := json.NewEncoder(&reqBody).Encode(&req); err != nil {
		return err
	}
	// 创建一个请求
	r, err := http.NewRequest("POST", tokenUrl, &reqBody)
	if err != nil {
		return fmt.Errorf("Error creating request: %w", err)
	}
	// 设置请求头
	r.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("Error sending request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	var res tokenRes
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return fmt.Errorf("Error reading response: %w", err)
	}
	*ak = res.AccessToken
	expiresAt := startTime.Add(time.Duration(res.ExpiresIn) * time.Second)
	stateMap[c.appid] = tokenState{res.AccessToken, expiresAt}
	return nil
}
