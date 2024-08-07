package tc_sms

import (
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

// 参考了GitHub上的代码
// https://github.com/ixre/go2o/blob/2c7f7c875501432b008b84636ab41cdac5527bd1/core/sp/tencent/tecent_sms.go
func Send(signName, templateId string, params, phones []string) (*sms.SendSmsResponse, error) {
	// 硬编码密钥到代码中有可能随代码泄露而暴露，有安全隐患，并不推荐。
	// 为了保护密钥安全，建议将密钥设置在环境变量中或者配置文件中，请参考本文凭证管理章节。
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1"
	client, _ := sms.NewClient(credential, regions.Guangzhou, cpf)

	req := sms.NewSendSmsRequest()
	// 配置签名和应用Id
	req.SmsSdkAppId = common.StringPtr(os.Getenv("TC_SDK_APP_ID"))
	req.SignName = common.StringPtr(signName)
	req.TemplateId = common.StringPtr(templateId)
	req.TemplateParamSet = common.StringPtrs(params)
	req.PhoneNumberSet = common.StringPtrs(phones)
	return client.SendSms(req)
}
