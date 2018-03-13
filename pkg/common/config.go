package common

import (
	"encoding/json"
	"log"
	"os"
)

// Config returns configuration parameters about database connection and credentials
type Config struct {
	PgHost        string
	PgPort        string
	PgUser        string
	PgPass        string
	DbName        string
	MongoHost     string
	MongoPort     string
	MongoUser     string
	MongoPassword string
}

// LoadConfiguration retrives configuration information from file and returns a setting struct
func LoadConfiguration(configurationFilePath string) Config {
	file, err := os.Open(configurationFilePath)
	if err != nil {
		log.Panic("Unable to open configuration file")
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Config{}

	error := decoder.Decode(&conf)
	if error != nil {
		log.Panic("Error decoding configuration file")
	}

	return conf
}
