package database

import "github.com/jinzhu/gorm"

var (
	// Session provides gloabaly available database connection
	Session *gorm.DB
)
