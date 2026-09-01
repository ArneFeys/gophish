package api

import (
	"net/http"
	"os/exec"
	"regexp"

	log "github.com/gophish/gophish/logger"
	"github.com/gophish/gophish/models"
)

// hostPattern matches a plain DNS name and nothing else, so no shell
// metacharacter can reach the resolver command.
var hostPattern = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9-]{0,61}[A-Za-z0-9])?(\.[A-Za-z0-9]([A-Za-z0-9-]{0,61}[A-Za-z0-9])?)*$`)

// DNSCheck resolves the host given in the "host" query parameter and returns
// the raw resolver output, so an operator can debug a sending domain from the
// UI instead of shelling into the box.
func (as *Server) DNSCheck(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	if host == "" {
		JSONResponse(w, models.Response{Success: false, Message: "No host given"}, http.StatusBadRequest)
		return
	}
	if len(host) > 253 || !hostPattern.MatchString(host) {
		JSONResponse(w, models.Response{Success: false, Message: "Invalid host"}, http.StatusBadRequest)
		return
	}
	out, err := exec.Command("getent", "hosts", host).CombinedOutput()
	if err != nil {
		log.Error(err)
	}
	JSONResponse(w, models.Response{Success: true, Message: string(out)}, http.StatusOK)
}
