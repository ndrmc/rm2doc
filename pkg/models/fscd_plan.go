package models

// FscdPlan represents food requirement for PSNP program
type FscdPlan struct {
	Base
	ID       int64
	Name     string
	Code     string
	Year     string
	Duration int32
	Status   int32
}
