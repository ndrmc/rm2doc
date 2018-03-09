package models

import (
	"time"
)

// BaseModel contains fields and methods common to all transactional entities
type BaseModel struct {
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedBy  int       `json:"created_by"`
	ModifiedBy int       `json:"modified_by"`
	DeletedAt  string    `json:"deleted_at"`
}

// Hash generates a hashed value of the object
func (b *BaseModel) Hash() string {
	return ""
}
