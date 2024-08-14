package bec_http

import (
	"encoding/json"
	"net/http"
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
func SendResBytes(w http.ResponseWriter, bytes []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func SendErr(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if e, ok := err.(IStatusError); ok {
		status = e.Status()
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status": status,
		"error":  err.Error(),
	})
}
func SendRes(w http.ResponseWriter, res any) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if res == nil {
		res = EmptyRes
	}
	json.NewEncoder(w).Encode(res)
}
