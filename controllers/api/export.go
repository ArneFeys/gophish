package api

import (
	"net/http"
	"os"
	"path/filepath"

	log "github.com/gophish/gophish/logger"
	"github.com/gophish/gophish/models"
)

// reportDir is where generated campaign reports are written.
const reportDir = "static/reports"

// ReportDownload streams a previously generated campaign report back to the
// caller. The report is picked with the "name" query string parameter.
func (as *Server) ReportDownload(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		JSONResponse(w, models.Response{Success: false, Message: "No report name given"}, http.StatusBadRequest)
		return
	}
	// Only a bare file name is accepted, so the download stays inside
	// reportDir no matter what traversal the caller asks for.
	name = filepath.Base(filepath.Clean("/" + name))
	data, err := os.ReadFile(filepath.Join(reportDir, name))
	if err != nil {
		log.Error(err)
		JSONResponse(w, models.Response{Success: false, Message: "Report not found"}, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+name)
	w.Write(data)
}
