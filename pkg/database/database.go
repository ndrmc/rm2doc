package database

import (
	"database/sql"
	//Importing postgres library
)

var (
	// Con provides gloabaly available database connection
	Con *sql.DB
)
