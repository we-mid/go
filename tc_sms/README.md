# tc_sms

> 简单易用的 腾讯云SMS短信发送 SDK

### 基础示例

```ini
# .env
TENCENTCLOUD_SECRET_ID=AKIDxxxxxx
TENCENTCLOUD_SECRET_KEY=APxxxxxx
TC_SDK_APP_ID=1xxxxxx
```

```go
// sms_job.go
import (
  // 自动加载.env环境变量
	_ "github.com/joho/godotenv/autoload"
	// 引入tc_sms
	"gitee.com/we-mid/go/tc_sms"
)

var (
	// 腾讯云-短信-国内短信-正文模板管理
	// https://console.cloud.tencent.com/smsv2/csms-template
	templateId = "2197493" // 运营数据通知
	template = `{1} 近期运营数据:
*{2} PV:[{3}] UV:[{4}] 留存:[{5}]
*{6} PV:[{7}] UV:[{8}] 留存:[{9}]
*{10} PV:[{11}] UV:[{12}] 留存:[{13}]
*{14} PV:[{15}] UV:[{16}] 留存:[{17}]`
	feeLimit = 3 // 单条消息发送fee上限设定
	sigName = "XX应用" // 短信签名
	phones = []string{"+8613xxxx", "+8618xxxx", "+8613xxxx"} // 接收手机号码
)

func smsJob() {
	// 组装参数
	var params []string
	// ...
	// 发送前预览完整文本，自动计算fee
	result, n, fee := tc_sms.SmsPreview(template, params)
	log.Println("smsJob preview:", result)
	log.Printf("smsJob preview: n=%d, fee=%d\n", n, fee)
	// 针对fee设置风控门槛
	if fee > feeLimit {
		log.Printf("smsJob fee exceeded! n=%d, fee=%d\n", n, fee)
		return
	}
	// 发送短信
	res, err := tc_sms.SmsSend(signName, templateId, params, phones)
	if err != nil {
		log.Println("smsJob error:", err)
		return
	}
	// 检查返回结果是否全部发送成功
	str := res.ToJsonString()
	if strings.Count(str, "send success") < len(phones) {
		log.Println("smsJob some failed", str)
		return
	}
	log.Println("smsJob all success:", str)
}
```

### 错误码 InvalidParameterValue.TemplateParameterLengthLimit 及文本截断

```go
// 注意：错误码 InvalidParameterValue.TemplateParameterLengthLimit
// 非验证码短信：每个变量取值最多支持6个字符。
// https://cloud.tencent.cn/document/product/382/52075
const paramLimit = 6

func smsJob() {
	// ...
	for i, v := range params {
		// ❌ 注意：错误的截断方式
		// if len(v) > paramLimit {
		// 	params[i] = v[:paramLimit]
		// }
		// ✅ 注意：正确处理中文字符长度
		if len([]rune(v)) > paramLimit {
			params[i] = truncateString(v, paramLimit) // 直接截断
			// ...或进行其他自定义的智能截断处理
		}
	}
	// ...
}

func truncateString(s string, length int) string {
	runes := []rune(s)
	if len(runes) <= length {
		return s
	}
	return string(runes[:length])
}
```

### 参考链接

- 腾讯云-短信-国内短信-正文模板管理：https://console.cloud.tencent.com/smsv2/csms-template
- 关于单条短信计费fee：https://console.cloud.tencent.com/smsv2/csms-template/create
- 非验证码短信：每个变量取值最多支持6个字符：https://cloud.tencent.cn/document/product/382/52075
- 参考了GitHub上的代码：https://github.com/ixre/go2o/blob/2c7f7c875501432b008b84636ab41cdac5527bd1/core/sp/tencent/tecent_sms.go
