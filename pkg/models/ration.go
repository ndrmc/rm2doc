package models

// Ration is used to define the amount of assistance given to a beneficiary in a given operation period
type Ration struct {
	Base
	ID          int64
	ReferenceNo string
	Description string
	Current     bool
}
