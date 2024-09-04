# basic_auth

> HTTP Basic-Auth 辅助函数库

```ini
# .env
BASICAUTH_USERLIST='[{"user":"xxx","pass":"xxxxxx"},{"user":"xxx","pass":"xxxxxx"}]'
```

```go
import (
	// ...
	"gitee.com/we-mid/go/basic_auth"
	"gitee.com/we-mid/go/util"
)

func init() {
	if err := basic_auth.InitFromEnv(); err != nil {
		log.Fatalln("basic_auth.InitFromEnv:", err)
	}
}

func main() {
	// API路由
	http.HandleFunc("/api/foo", util.HandlerWrap(apiFoo))
	http.HandleFunc("/api/bar", util.HandlerWrap(apiBar))
	// 页面路由，创建一个文件服务器
	fs := http.FileServer(http.Dir("./public"))
	http.HandleFunc("/", basic_auth.Wrap(fs.ServeHTTP))
	// ...
}

func apiFoo(w http.ResponseWriter, r *http.Request) error {
	if err := util.EnableCors(w, r); err != nil {
		return err
	}
	if user := basic_auth.Check(r); user == "" {
		return util.Err401
	}
	// ...
}
```
