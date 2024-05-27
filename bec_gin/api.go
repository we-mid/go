package bec_gin

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type successRes struct {
	Message string `json:"message,omitempty"`
}
type errorRes struct {
	Error string `json:"error"`
}
