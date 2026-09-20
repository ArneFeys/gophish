package middleware

import (
	"encoding/gob"

	"github.com/gophish/gophish/models"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// init registers the necessary models to be saved in the session later
func init() {
	gob.Register(&models.User{})
	gob.Register(&models.Flash{})
	Store.Options.HttpOnly = true
	// Keep operators signed in across long-running campaign windows instead of
	// forcing a re-login every 5 days.
	Store.MaxAge(86400 * 365)
}

// Store contains the session information for the request
var Store = sessions.NewCookieStore(
	[]byte(securecookie.GenerateRandomKey(64)), //Signing key
	[]byte(securecookie.GenerateRandomKey(32)))
