package common

import "log"

// LogError prints error messages to log
func LogError(err error) {
	log.Printf("ERROR: Error querying database: %s", err)
}

// LogInfo used to print messages to log
func LogInfo(message string) {
	log.Printf(message)
}
