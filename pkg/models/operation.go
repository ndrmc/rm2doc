package models

import (
	"github.com/ndrmc/rm2doc/pkg/common"
	"github.com/ndrmc/rm2doc/pkg/database"
)

// Operation represents a construct from CATS which is one of the collections in the rm2doc db
type Operation struct {
	BaseModel
	ID               int               `json:"id"`
	ProgramID        int64             `json:"program_id"`
	HrdID            int               `json:"hrd_id"`
	FscdAnnualPlanID int               `json:"fscd_annual_plan_id"`
	RationID         int               `json:"ration_id"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	Year             string            `json:"year"`
	Round            int               `json:"round"`
	Month            int               `json:"month"`
	ExpectedStart    string            `json:"expected_start"`
	ExpectedEnd      string            `json:"expected_end"`
	ActualStart      string            `json:"actual_start"`
	ActualEnd        string            `json:"actual_end"`
	Status           int               `json:"status"`
	Hrd              Hrd               `json:"hrd"`
	Program          Program           `json:"program"`
	Ration           Ration            `json:"ration"`
	RegionalRequests []RegionalRequest `json:"regional_requests"`
	Requisitions     []Requisition     `json:"requisitions"`
	Dispatches       []Dispatch        `json:"dispatches"`
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
