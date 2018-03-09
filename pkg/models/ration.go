package models

// Ration is used to define the amount of assistance given to a beneficiary in a given operation period
type Ration struct {
	BaseModel
	ID          int    `json:"id"`
	ReferenceNo string `json:"reference_no"`
	Description string `json:"description"`
	Current     bool   `json:"current"`
}
