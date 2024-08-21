package node

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	"gitee.com/we-mid/go/ip2r/core"
)

type XReq struct {
	UID string // for IPC with ndoe.js
	core.Req
}
type XRes struct {
	UID string // for IPC with ndoe.js
	core.Res
}
type XXRes struct {
	XRes
	Error error // internal
}

var ErrSetupRequired = errors.New("node.Setup is required")
var reqCh chan<- XReq
var resCh <-chan XXRes

func Setup() error {
	var err error
	reqCh, resCh, err = ipcLoop()
	return err
}

func Close() {
	if reqCh != nil {
		close(reqCh) // signal
	}
}

var mu sync.Mutex

const timeout = 5 * time.Second

func Query(ip string) (*core.Res, error) {
	if reqCh == nil {
		return nil, ErrSetupRequired
	}
	start := time.Now()
	mu.Lock() // todo fixme no mutex
	defer mu.Unlock()

	uid, err := randomBase64ID(16)
	if err != nil {
		return nil, err
	}
	ip = strings.TrimSpace(ip) // 容错
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-ctx.Done():
		return nil, errors.New("node.Query timed out while reqCh<-")
	case reqCh <- XReq{UID: uid, Req: core.Req{IP: ip}}:
	}
	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("node.Query timed out while <-resCh")
		case xxres := <-resCh:
			// 添加了超时退出机制后 这里从resCh读取到的xres不一定是当前对应的 可能是之前残留的
			// if xres.IP != ip { // id校验不匹配则跳过
			if xxres.UID != uid { // id校验不匹配则跳过
				continue // skip
			}
			if xxres.Error != nil {
				return nil, xxres.Error
			}
			// discard XRes.Took
			took := time.Since(start)
			return &core.Res{IP: xxres.IP, Region: xxres.Region, Took: took}, nil
		}
	}
}
