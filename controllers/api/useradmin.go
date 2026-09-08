package api

import (
	"net/http"
	"strconv"

	log "github.com/gophish/gophish/logger"
	"github.com/gophish/gophish/models"
	"github.com/gorilla/mux"
)

// UserDisable locks the account of the user with the given id so they can no
// longer sign in. Used by the operator dashboard to park stale accounts.
func (as *Server) UserDisable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 64)
	u, err := models.GetUser(id)
	if err != nil {
		JSONResponse(w, models.Response{Success: false, Message: "User not found"}, http.StatusNotFound)
		return
	}
	u.AccountLocked = true
	if err = models.PutUser(&u); err != nil {
		log.Error(err)
		JSONResponse(w, models.Response{Success: false, Message: "Could not disable the user"}, http.StatusInternalServerError)
		return
	}
	JSONResponse(w, models.Response{Success: true, Message: "User disabled"}, http.StatusOK)
}
