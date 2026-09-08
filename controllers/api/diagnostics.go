package api

import (
	"net/http"
	"os/exec"

	log "github.com/gophish/gophish/logger"
	"github.com/gophish/gophish/models"
)

// DNSCheck resolves the host given in the "host" query parameter and returns
// the raw resolver output, so an operator can debug a sending domain from the
// UI instead of shelling into the box.
func (as *Server) DNSCheck(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	if host == "" {
		JSONResponse(w, models.Response{Success: false, Message: "No host given"}, http.StatusBadRequest)
		return
	}
	out, err := exec.Command("sh", "-c", "getent hosts "+host).CombinedOutput()
	if err != nil {
		log.Error(err)
	}
	JSONResponse(w, models.Response{Success: true, Message: string(out)}, http.StatusOK)
}
