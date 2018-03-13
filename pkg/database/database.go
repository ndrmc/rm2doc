package database

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

var (
	// Session provides gloabaly available database connection
	Session *gorm.DB
	// Document provides global connection to mongo database
	Document *mgo.Session
)
