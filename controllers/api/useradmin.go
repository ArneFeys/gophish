package api

import (
	"net/http"
	"strconv"

	ctx "github.com/gophish/gophish/context"
	log "github.com/gophish/gophish/logger"
	"github.com/gophish/gophish/models"
	"github.com/gorilla/mux"
)

// UserDisable locks the account of the user with the given id so they can no
// longer sign in. Only users with the ModifySystem permission may do this.
func (as *Server) UserDisable(w http.ResponseWriter, r *http.Request) {
	currentUser := ctx.Get(r, "user").(models.User)
	hasSystem, err := currentUser.HasPermission(models.PermissionModifySystem)
	if err != nil {
		JSONResponse(w, models.Response{Success: false, Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	if !hasSystem {
		JSONResponse(w, models.Response{Success: false, Message: http.StatusText(http.StatusForbidden)}, http.StatusForbidden)
		return
	}
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
