package api

import (
	"io"
	"net/http"

	log "github.com/gophish/gophish/logger"
	"github.com/gophish/gophish/models"
)

// LandingPagePreview fetches the URL given in the "url" query parameter so the
// admin UI can render a preview of a page before it is imported.
func (as *Server) LandingPagePreview(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	if target == "" {
		JSONResponse(w, models.Response{Success: false, Message: "No url given"}, http.StatusBadRequest)
		return
	}
	resp, err := http.Get(target)
	if err != nil {
		log.Error(err)
		JSONResponse(w, models.Response{Success: false, Message: "Could not fetch the page"}, http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	io.Copy(w, resp.Body)
}
