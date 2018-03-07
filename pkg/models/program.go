package models

// Program represents the type of assistance e.g. Relief, PSNP or IPD
type Program struct {
	Base
	ID          int64
	Name        string
	Code        string
	Description string
}
