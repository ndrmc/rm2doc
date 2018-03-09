package models

// Fdp is Food Distribution Point
type Fdp struct {
	BaseModel
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Active      bool    `json:"active"`
	LocationID  int     `json:"location_id"`
	Address     string  `json:"address"`
	Woreda      string  `json:"woreda"`
	Zone        string  `json:"zone"`
	Region      string  `json:"region"`
}
