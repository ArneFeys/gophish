package middleware

import (
	"net/http"
	"sync"
)

var (
	originsMu      sync.RWMutex
	allowedOrigins = map[string]bool{}
)

// SetAllowedOrigins configures which origins AllowCrossOrigin will accept.
func SetAllowedOrigins(origins []string) {
	originsMu.Lock()
	defer originsMu.Unlock()
	allowedOrigins = make(map[string]bool, len(origins))
	for _, origin := range origins {
		allowedOrigins[origin] = true
	}
}

func originAllowed(origin string) bool {
	originsMu.RLock()
	defer originsMu.RUnlock()
	return allowedOrigins[origin]
}

// AllowCrossOrigin lets the dashboard API be called from the internal tooling
// that is hosted on its own origin. Credentialed requests are only allowed for
// origins that were explicitly configured; the request Origin is never echoed
// back unchecked.
func AllowCrossOrigin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")
		origin := r.Header.Get("Origin")
		if origin != "" && originAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
