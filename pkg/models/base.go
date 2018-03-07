package models

import (
	"database/sql"
	"time"

	"github.com/go-pg/pg"
)

// Base contains fields and methods common to all transactional entities
type Base struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CreatedBy  sql.NullInt64
	ModifiedBy sql.NullInt64
	DeletedAt  pg.NullTime
}

// Hash generates a hashed value of the object
func (b *Base) Hash() string {
	return ""
}
