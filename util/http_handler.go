package util

import (
	"errors"
	"net/http"
	"time"
)

func HandlerWrap(
	logic func(w http.ResponseWriter, r *http.Request) (any, error),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		res, err := logic(w, r)
		if err != nil {
			if errors.Is(err, ErrHandledAndBreak) {
				HttpLog(r, start, 200, err)
			} else {
				if e, ok := err.(*StatusError); ok && e.Status < 500 {
					SendErr(w, err)
					HttpLog(r, start, e.Status, err)
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
