package models

import (
	"github.com/ndrmc/rm2doc/pkg/database"
)

// Hrd represents food requirement for relief program
type Hrd struct {
	BaseModel
	ID        int `json:"id" bson:"_id"`
	YearGC    int `json:"year_gc" bson:"year_gc"`
	Status    int `json:"status" bson:"status"`
	YearEC    int `json:"year_ec" bson:"year_ec"`
	MonthFrom int `json:"month_from" bson:"month_from"`
	Duration  int `json:"duration" bson:"duration"`
	SeasonID  int `json:"season_id" bson:"season_id"`
	RationID  int `json:"ration_id" bson:"ration_id"`
}

// GetHrd returns an Hrd record from transactional database
func GetHrd(id int) Hrd {
	var hrd Hrd
	database.Session.Find(&hrd, id)
	return hrd
}
