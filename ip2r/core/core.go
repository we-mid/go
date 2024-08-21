package core

import "time"

type Req struct {
	IP string
	// 更多控制参数可拓展
}
type Res struct {
	IP     string
	Region string
	Took   time.Duration
	// todo more
}
