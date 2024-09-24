package passwordless

import (
	"net/http"

	"gitee.com/we-mid/go/bec_http"
)

func (p *Passwordless) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(p.RoutePrefix+"/attempt", bec_http.HandlerWrap(p.HandleAttempt))
	mux.HandleFunc(p.RoutePrefix+"/verify", bec_http.HandlerWrap(p.HandleVerify))
	mux.HandleFunc(p.RoutePrefix+"/session", bec_http.HandlerWrap(p.HandleSession))
}
