package api

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gophish/gophish/models"
)

// CheckHostConnectivity pings a host to verify network connectivity
// for SMTP troubleshooting.
func (as *Server) CheckHostConnectivity(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	if host == "" {
		JSONResponse(w, models.Response{Success: false, Message: "host parameter required"}, http.StatusBadRequest)
		return
	}

	cmd := exec.Command("sh", "-c", fmt.Sprintf("ping -c 3 %s", host))
	output, err := cmd.CombinedOutput()
	if err != nil {
		JSONResponse(w, models.Response{Success: false, Message: string(output)}, http.StatusInternalServerError)
		return
	}

	JSONResponse(w, map[string]string{"output": string(output)}, http.StatusOK)
}
