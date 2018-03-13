package models

// Ration is used to define the amount of assistance given to a beneficiary in a given operation period
type Ration struct {
	BaseModel
	ID          int    `json:"id" bson:"_id"`
	ReferenceNo string `json:"reference_no" bson:"reference_no"`
	Description string `json:"description" bson:"description"`
	Current     bool   `json:"current" bson:"current"`
}
