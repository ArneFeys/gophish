package api

import (
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"

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
	if err := validatePreviewURL(target); err != nil {
		JSONResponse(w, models.Response{Success: false, Message: err.Error()}, http.StatusBadRequest)
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

// validatePreviewURL rejects anything that is not a plain http(s) URL resolving
// to a public address, so the preview cannot be pointed at internal services or
// the cloud metadata endpoint.
func validatePreviewURL(raw string) error {
	u, err := url.Parse(raw)
	if err != nil {
		return errors.New("Invalid url")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("Only http and https urls can be previewed")
	}
	host := u.Hostname()
	if host == "" {
		return errors.New("Invalid url")
	}
	ips, err := net.LookupIP(host)
	if err != nil || len(ips) == 0 {
		return errors.New("Could not resolve the host")
	}
	for _, ip := range ips {
		if ip.IsLoopback() || ip.IsPrivate() || ip.IsLinkLocalUnicast() ||
			ip.IsLinkLocalMulticast() || ip.IsUnspecified() {
			return errors.New("Refusing to fetch an internal address")
		}
	}
	return nil
}
