# mpwx

> 微信公众号-模板消息推送 Go语言SDK

**基础用法：固定模板，推送至固定粉丝用户，适用于管理员消息推送**

```
# .env
WXPUSH_APPID=wxxxxxxxxxxxxxxxxx
WXPUSH_SECRET=00xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
WXPUSH_TEMPLATE_PLAIN=ZZxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
WXPUSH_USER_ADMIN=onxxxxxxxxxxxxxxxxxxxxxxxxxx
```

注：此固定模板内容必须设置为：`> {{text.DATA}}`，注意开头必须追加某些非空字符，否则消息无法正常展示

```go
import "gitee.com/we-mid/go/mpwx"

// 环境变量默认前缀为"WXPUSH"=>WXPUSH_APPID...
mpwx.Default.PushToAdmin("Hello")
mpwx.Default.PushToAdminf("%q %v", city, err)
mpwx.Default.PushToAdminf("%q-%q added", prov, city)

// 也可以通过指定不同的前缀，创建多个实例
mp1 := mpwx.NewFromEnv("MP1") // MP1_APPID...
mp2 := mpwx.NewFromEnv("MP2") // MP2_APPID...
mp1.PushToAdminf("%q %v", city, err)
mp2.PushToAdminf("%q-%q added", prov, city)
```

**高级用法：任意模板，推送至任意粉丝用户，适用于更多场景**

```go
touser := "wxxxxxxxxxxxxxxxxx"
template_id := "ZZxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
tMap := map[string]string{
	"key1": "value1",
	"key2": "value2",
}
err := mp1.SendTemplateMessage(touser, template_id, tMap)
```

**注意事项：关于协程**

1. `PushToAdmin`系列函数默认调用协程并打印错误，如希望阻塞并返回错误则应调用`PushToAdminSync`
2. 由于调用了协程，因此在main函数优雅退出前可以调用`mpwx.WG.Wait`，等待相关协程完成

**命令行工具**

```sh
go install gitee.com/we-mid/go/mpwx/cmd/wxpush@latest
ls .env  # 指定相关环境变量
wxpush "Hello World"
```

**相关资料**

微信公众号文档：基础消息能力 /模板消息接口
https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html

微信公众平台接口测试帐号申请
无需公众帐号、快速申请接口测试号，直接体验和测试公众平台所有高级接口
https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login
请注意：
1、测试模板的模板ID仅用于测试，不能用来给正式帐号发送模板消息
2、为方便测试，测试模板可任意指定内容，但实际上正式帐号的模板消息，只能从模板库中获得
3、需为正式帐号申请新增符合要求的模板，需使用正式号登录公众平台，按指引申请
4、模板内容可设置参数(模板标题不可)，供接口调用时使用，参数需以{{开头，以.DATA}}结尾
