package models

import (
	"github.com/ndrmc/rm2doc/pkg/common"
	"github.com/ndrmc/rm2doc/pkg/database"
	"gopkg.in/mgo.v2/bson"
)

// Operation represents a construct from CATS which is one of the collections in the rm2doc db
type Operation struct {
	BaseModel
	ID               int               `json:"id" bson:"_id"`
	ProgramID        int64             `json:"program_id" bson:"program_id"`
	HrdID            int               `json:"hrd_id" bson:"hrd_id"`
	FscdAnnualPlanID int               `json:"fscd_annual_plan_id" bson:"fscd_annual_plan_id"`
	RationID         int               `json:"ration_id" bson:"ration_id"`
	Name             string            `json:"name" bson:"name"`
	Description      string            `json:"description" bson:"description"`
	Year             string            `json:"year" bson:"year"`
	Round            int               `json:"round" bson:"round"`
	Month            int               `json:"month" bson:"month"`
	ExpectedStart    string            `json:"expected_start" bson:"expected_start"`
	ExpectedEnd      string            `json:"expected_end" bson:"expected_end"`
	ActualStart      string            `json:"actual_start" bson:"actual_start"`
	ActualEnd        string            `json:"actual_end" bson:"actual_end"`
	Status           int               `json:"status" bson:"status"`
	Hrd              Hrd               `json:"hrd" bson:"hrd"`
	Program          Program           `json:"program" bson:"program"`
	Ration           Ration            `json:"ration" bson:"ration"`
	RegionalRequests []RegionalRequest `json:"regional_requests" bson:"regional_requests"`
	Requisitions     []Requisition     `json:"requisitions" bson:"requisitions"`
	Dispatches       []Dispatch        `json:"dispatches" bson:"dispatch"`
}

// GetOperation returns an operation record from transactional database
func GetOperation(id int) Operation {
	var operation Operation
	errs := database.Session.First(&operation, id).GetErrors()

	for _, err := range errs {
		common.LogError(err)
	}
	return operation
}

// GetOperationGraph returns an operation record and all associated relations from transactional database
func GetOperationGraph(id int) Operation {
	var operation Operation
	errs := database.Session.
		Preload("Hrd").
		Preload("Program").
		Preload("Ration").
		Preload("Dispatches").
		Preload("Dispatches.Items").
		Preload("Requisitions").
		Preload("Requisitions.Items").
		Preload("Requisitions.Items.Fdp").
		Preload("RegionalRequests").
		Preload("RegionalRequests.Items").
		First(&operation, id).GetErrors()

	for _, err := range errs {
		common.LogError(err)
	}
	return operation
}

// GetAllOperations returns all operations in the database
func GetAllOperations() []Operation {
	var operations []Operation
	database.Session.Find(&operations)
	return operations
}

// HowManyOperations returnes the total number of operations in the database
func HowManyOperations() int {
	var count int
	database.Session.Model(&Operation{}).Count(&count)
	return count
}

// UpdateOperationDocument creates corresponding records from relation model in mongo
func UpdateOperationDocument(operation Operation) {
	c := database.Document.DB("cats_analytics").C("operations")
	err := c.Insert(&operation)

	if err != nil {
		common.LogError(err)
	}
}

// GetOperationDocument returns graph of an operation record
func GetOperationDocument(id int) *Operation {
	var operation *Operation
	if err := database.Document.DB("cats_analytics").C("operations").Find(bson.M{"_id": id}).One(&operation); err != nil {
		common.LogError(err)
	}

	return operation
}
