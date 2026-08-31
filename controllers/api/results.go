package api

import (
	"net/http"
	"strconv"

	ctx "github.com/gophish/gophish/context"
	log "github.com/gophish/gophish/logger"
	"github.com/gophish/gophish/models"
	"github.com/gorilla/mux"
)

// CampaignResultsSorted returns the results of a campaign ordered by the column
// named in the "sort" query string parameter.
func (as *Server) CampaignResultsSorted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 64)
	sortBy := r.URL.Query().Get("sort")
	if sortBy == "" {
		sortBy = "send_date"
	}
	rs, err := models.ResultsSorted(id, ctx.Get(r, "user_id").(int64), sortBy)
	if err != nil {
		log.Error(err)
		JSONResponse(w, models.Response{Success: false, Message: "Could not load the results"}, http.StatusInternalServerError)
		return
	}
	JSONResponse(w, rs, http.StatusOK)
}
