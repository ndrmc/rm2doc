package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/ndrmc/rm2doc/pkg/common"
	"github.com/ndrmc/rm2doc/pkg/database"
	"github.com/ndrmc/rm2doc/pkg/models"
	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// FindOperation fetches operation record from cats_analytics db
func FindOperation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	operation := models.GetOperationDocument(id)
	respondWithJson(w, http.StatusOK, operation)
}

func main() {
	conf := common.LoadConfiguration("/Users/yared/src/gospace/src/github.com/ndrmc/rm2doc/config.json")
	initDB(conf)

	r := mux.NewRouter()
	r.HandleFunc("/operations/{id}", FindOperation).Methods("GET")

	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func initDB(conf common.Config) {
	var pgErr, mongoErr error

	pgInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.PgHost, conf.PgPort, conf.PgUser, conf.PgPass, conf.DbName)
	database.Session, pgErr = gorm.Open("postgres", pgInfo)
	if pgErr != nil {
		log.Panicf("Error making connection to database. Detail: %s", pgErr)
	}

	common.LogInfo("Successfuly connected to database")

	// mongoInfo := fmt.Sprintf("mongodb://%s:%s@%s:%s", conf.MongoUser, conf.MongoPassword, conf.MongoHost, conf.MongoPort)
	database.Document, mongoErr = mgo.Dial("localhost")
	if mongoErr != nil {
		common.LogError(mongoErr)
	}

	mongoErr = database.Document.Ping()
	if mongoErr != nil {
		common.LogError(mongoErr)
	}

}
