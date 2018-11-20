package models

// Program represents the type of assistance e.g. Relief, PSNP or IPD
type Program struct {
	BaseModel
	ID          int    `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Code        string `json:"code" bson:"code"`
	Description string `json:"description" bson:"description"`
}
