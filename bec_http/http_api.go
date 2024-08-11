package bec_http

import (
	"encoding/json"
	"net/http"
)

var (
	Err400 = NewStatusError(400, nil)
	Err401 = NewStatusError(401, nil)
	Err404 = NewStatusError(404, nil)
	Err405 = NewStatusError(405, nil)
	Err429 = NewStatusError(429, nil)
	Err500 = NewStatusError(500, nil)

	EmptyRes = struct{}{}
)

func SendErrText(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if e, ok := err.(*StatusError); ok {
		status = e.Status()
	}
	http.Error(w, err.Error(), status)
}
func SendResText(w http.ResponseWriter, res string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func SendErr(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if e, ok := err.(IStatusError); ok {
		status = e.Status()
	}
	w.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(map[string]any{
		"status": status,
		"error":  err.Error(),
	})
	http.Error(w, string(bytes), status)
}
func SendRes(w http.ResponseWriter, res any) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if res == nil {
		res = EmptyRes
	}
	bytes, _ := json.Marshal(res)
	w.Write(bytes)
}
