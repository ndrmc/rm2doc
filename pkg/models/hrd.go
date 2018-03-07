package models

import (
	"database/sql"

	"github.com/ndrmc/analytics/pkg/common"
	"github.com/ndrmc/analytics/pkg/database"
)

// Hrd represents food requirement for relief program
type Hrd struct {
	Base
	ID        int64         `json:"id"`
	YearGC    int32         `json:"year_gc"`
	Status    int32         `json:"status"`
	YearEC    sql.NullInt64 `json:"year_ec"`
	MonthFrom sql.NullInt64 `json:"month_from"`
	Duration  sql.NullInt64 `json:"duration"`
	SeasonID  int64         `json:"season_id"`
	RationID  int64         `json:"ration_id"`
}

// GetHrd returns an Hrd record from transactional database
func GetHrd(id int64) *Hrd {
	var stmt = "select * from hrds where id=$1"
	rows, err := database.Con.Query(stmt, id)
	if err != nil {
		common.LogError(err)
	}

	hrds, err := mapHrds(rows)
	if err != nil {
		common.LogError(err)
	}

	return hrds[0]
}

func mapHrds(rows *sql.Rows) ([]*Hrd, error) {
	var err error
	hrds := make([]*Hrd, 0)

	for rows.Next() {
		h := new(Hrd)
		err = rows.Scan(
			&h.ID,
			&h.YearGC,
			&h.Status,
			&h.MonthFrom,
			&h.Duration,
			&h.SeasonID,
			&h.RationID,
			&h.CreatedAt,
			&h.UpdatedAt,
			&h.CreatedBy,
			&h.ModifiedBy,
			&h.DeletedAt,
			&h.YearEC)

		if err != nil {
			panic(err)
		}
		hrds = append(hrds, h)
	}

	return hrds, err
}
