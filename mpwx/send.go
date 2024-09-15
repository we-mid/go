package mpwx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"gitee.com/we-mid/go/util"
)

const sendUrl = "https://api.weixin.qq.com/cgi-bin/message/template/send"

type rawValueItem struct {
	Value string `json:"value"`
}
type rawValueMap map[string]rawValueItem

type sendReq struct {
	ToUser     string      `json:"touser"`
	TemplateId string      `json:"template_id"`
	Data       rawValueMap `json:"data"`
}

type sendRes struct {
	ErrCode int `json:"errcode"`
}

// 基础消息能力 /模板消息接口
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (c *MpwxClient) SendTemplateMessage(touser, template_id string, valueMap map[string]string) error {
	// 设置url+query
	var ak string
	if err := c.accessToken(&ak); err != nil {
		return err
	}
	query := url.Values{}
	query.Set("access_token", ak)
	u, err := util.URLParseQueryPatch(sendUrl, query)
	if err != nil {
		return err
	}
	// 设置请求体
	var reqBody bytes.Buffer
	rMap := make(rawValueMap)
	for k, v := range valueMap {
		rMap[k] = rawValueItem{v}
	}
	req := sendReq{touser, template_id, rMap}
	if err := json.NewEncoder(&reqBody).Encode(req); err != nil {
		return err
	}
	// 创建一个请求
	r, err := http.NewRequest("POST", u.String(), &reqBody)
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
	var res sendRes
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return fmt.Errorf("Error reading response: %w", err)
	}

	// Valid values of errCode
	// https://developers.weixin.qq.com/miniprogram/en/dev/api-backend/open-api/template-message/templateMessage.send.html
	// Interface error code
	// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/en/industry/express/express_open_msg.html#Interface%20error%20code
	// 40003: Error in openid parameter
	if res.ErrCode != 0 {
		return fmt.Errorf("sendRes=%+v", res)
	}
	return nil
}
