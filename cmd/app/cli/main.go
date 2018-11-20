package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/mgo.v2"

	"github.com/jinzhu/gorm"
	"github.com/ndrmc/rm2doc/pkg/database"

	_ "github.com/lib/pq"
	"github.com/ndrmc/rm2doc/pkg/common"
	"github.com/ndrmc/rm2doc/pkg/models"
)

func main() {
	conf := common.LoadConfiguration("./config.json")
	initDB(conf)
	// loadOperation(32)
	migrateOperations()
}

func countOperations() {
	count := models.HowManyOperations()
	fmt.Printf("Total number of Operations is: %d", count)
}

func loadOperation(id int) {
	result := models.GetOperationGraph(id)

	// models.UpdateOperationDocument(result)

	// fmt.Printf("Found operation %s and %d dispatches", result.Name, len(result.Dispatches))
	fmt.Printf("Operation is:  %s\n", result.Name)
	fmt.Printf("Total number of requistions: %d\n", len(result.Requisitions))
	fmt.Printf("Total number of dispatches: %d\n", len(result.Dispatches))

	// fmt.Printf("Total quantity in operation is:  %f", models.TotalDispatch(result.ID))
	fmt.Printf("Total number of beneficiaries in operation is: %f", models.TotalBeneficiaries(result.ID))
	fmt.Println("Saving operation metadata........")
	models.SaveOperationMetadata(result)

	// buf, err := json.Marshal(result)
	// if err != nil {
	// 	common.LogError(err)
	// }

	// fmt.Println(string(buf))
}

func migrateOperations() {
	all := models.GetAllOperations()

	for _, o := range all {
		value := models.GetOperationGraph(o.ID)
		models.SaveOperationMetadata(value)
		log.Printf("Migrated operation : %s", o.Name)
	}
}

func loadOperations() {
	total := models.GetAllOperations()
	// ops, err := json.MarshalIndent(total, "", "\t")
	ops, err := json.Marshal(total)
	if err != nil {
		common.LogError(err)
	}
	fmt.Println("===========================")
	fmt.Println(string(ops))
}

func initDB(conf common.Config) {
	var pgErr, mongoErr error

	pgInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.PgHost, conf.PgPort, conf.PgUser, conf.PgPass, conf.DbName)
	database.Session, pgErr = gorm.Open("postgres", pgInfo)
	if pgErr != nil {
		log.Panicf("Error making connection to database. Detail: %s", pgErr)
	}
	// database.Session.LogMode(true)
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
