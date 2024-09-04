package bec_http

import (
	"errors"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)
type Logic func(w http.ResponseWriter, r *http.Request) (any, error)
type LogicStream func(w http.ResponseWriter, r *http.Request, onBytes func([]byte)) error

var ErrHandledAndBreak = errors.New("handled and break")

var (
	Err400 = NewStatusError(400, nil)
	Err401 = NewStatusError(401, nil)
	Err403 = NewStatusError(403, nil)
	Err404 = NewStatusError(404, nil)
	Err405 = NewStatusError(405, nil)
	Err429 = NewStatusError(429, nil)
	Err500 = NewStatusError(500, nil)

	EmptyRes = struct{}{}
)
