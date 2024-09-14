package bec_http

import (
	"errors"
	"net/http"
	"time"
)

func CORSHandlerWrap[T any](logic Logic[T]) Handler {
	return HandlerWrap[T](func(w http.ResponseWriter, r *http.Request) (T, error) {
		if err := EnableCORS(w, r); err != nil {
			// fix: compiler: cannot use nil as T value in return statement
			// return nil, err
			// fix: compiler: invalid composite literal type T
			// return T{}, err
			var zero T
			return zero, err
		}
		return logic(w, r)
	})
}

func HandlerWrap[T any](logic Logic[T]) Handler {
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
