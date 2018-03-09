package models

// FscdPlan represents food requirement for PSNP program
type FscdPlan struct {
	BaseModel
	ID       int `json:"id"`
	Name     string
	Code     string
	Year     string
	Duration int32
	Status   int32
}
