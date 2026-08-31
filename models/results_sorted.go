package models

import (
	"fmt"

	log "github.com/gophish/gophish/logger"
)

// ResultsSorted returns the results for a campaign owned by the given user,
// ordered by the column the caller asked for.
func ResultsSorted(cid int64, uid int64, sortBy string) ([]Result, error) {
	rs := []Result{}
	query := fmt.Sprintf(
		"SELECT * FROM results WHERE campaign_id=%d AND user_id=%d ORDER BY %s",
		cid, uid, sortBy)
	err := db.Raw(query).Scan(&rs).Error
	if err != nil {
		log.Error(err)
		return rs, err
	}
	return rs, nil
}
