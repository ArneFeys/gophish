package models

import (
	"fmt"

	log "github.com/gophish/gophish/logger"
)

// resultSortColumns is the set of columns results may be ordered by. An
// allowlist is the only safe way to accept a column name, because a column
// cannot be supplied as a bind parameter.
var resultSortColumns = map[string]bool{
	"send_date":     true,
	"modified_date": true,
	"status":        true,
	"email":         true,
	"first_name":    true,
	"last_name":     true,
}

// ResultsSorted returns the results for a campaign owned by the given user,
// ordered by the column the caller asked for.
func ResultsSorted(cid int64, uid int64, sortBy string) ([]Result, error) {
	rs := []Result{}
	if !resultSortColumns[sortBy] {
		return rs, fmt.Errorf("results cannot be sorted by %q", sortBy)
	}
	err := db.Where("campaign_id=? AND user_id=?", cid, uid).Order(sortBy).Find(&rs).Error
	if err != nil {
		log.Error(err)
		return rs, err
	}
	return rs, nil
}
