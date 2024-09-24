package bec_http

import (
	"errors"
	"net/http"
	"time"

	"gitee.com/we-mid/go/util"
)

func CORSHandlerWrap[T any](logic Logic[T]) http.HandlerFunc {
	return HandlerWrap[T](func(w http.ResponseWriter, r *http.Request) (T, error) {
		if err := EnableCORS(w, r); err != nil {
			return util.ZeroValue[T](), err
		}
		return logic(w, r)
	})
}

func HandlerWrap[T any](logic Logic[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		res, err := logic(w, r)
		if err != nil {
			if errors.Is(err, ErrHandledAndBreak) {
				HttpLog(r, start, 200, err)
			} else {
				if e, ok := err.(IStatusError); ok && e.Status() < 500 {
					SendErr(w, err)
					HttpLog(r, start, e.Status(), err)
				} else {
					// protect error messages for 5xx
					SendErr(w, Err500)
					HttpLog(r, start, 500, err)
				}
			}
		} else {
			SendRes(w, res)
			HttpLog(r, start, 200, err)
		}
	}
}
