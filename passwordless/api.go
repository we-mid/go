package passwordless

import (
	"encoding/json"
	"net/http"

	"gitee.com/we-mid/go/bec_http"
)

func (p *Passwordless) handleReq(w http.ResponseWriter, r *http.Request, req any) error {
	if p.EnableCORS {
		if err := bec_http.EnableCORS(w, r); err != nil {
			return err
		}
	}
	if r.Method != http.MethodPost {
		// return bec_http.Err405
		return bec_http.Err400
	}
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		// return fmt.Errorf("json decode: %w", err)
		// return bec_http.NewStatusError(400, err)
		return bec_http.Err400
	}
	return nil
}
