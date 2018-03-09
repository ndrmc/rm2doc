package models

// Program represents the type of assistance e.g. Relief, PSNP or IPD
type Program struct {
	BaseModel
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
