package bec_http

import (
	"errors"
	"net/http"
	"time"
)

func CORSHandlerWrap(logic Logic) Handler {
	return HandlerWrap(func(w http.ResponseWriter, r *http.Request) (any, error) {
		if err := EnableCORS(w, r); err != nil {
			return nil, err
		}
		return logic(w, r)
	})
}

func HandlerWrap(logic Logic) Handler {
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