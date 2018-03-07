package models

import (
	"database/sql"

	"github.com/go-pg/pg"

	"github.com/ndrmc/analytics/pkg/database"

	"github.com/ndrmc/analytics/pkg/common"
)

// Operation represents a construct from CATS which is one of the collections in the analytics db
type Operation struct {
	Base
	ID               int64          `json:"id"`
	ProgramID        int64          `json:"program_id"`
	HrdID            sql.NullInt64  `json:"hrd_id"`
	FscdAnnualPlanID sql.NullInt64  `json:"fscd_annual_plan_id"`
	RationID         sql.NullInt64  `json:"ration_id"`
	Name             sql.NullString `json:"name"`
	Description      sql.NullString `json:"description"`
	Year             sql.NullString `json:"year"`
	Round            sql.NullInt64  `json:"round"`
	Month            sql.NullInt64  `json:"month"`
	ExpectedStart    pg.NullTime    `json:"expected_start"`
	ExpectedEnd      pg.NullTime    `json:"expected_end"`
	ActualStart      pg.NullTime    `json:"actual_start"`
	ActualEnd        pg.NullTime    `json:"actual_end"`
	Status           sql.NullInt64  `json:"status"`
	HrdInfo          *Hrd           `json:"hrd"`
	ProgramInfo      *Program       `json:"program"`
	RationInfo       *Ration        `json:"ration"`
	Dispatches       []*Dispatch    `json:"dispatches"`
}

// GetOperation returns an operation record from transactional database
func GetOperation(id int64) *Operation {
	var stmt = "select * from operations where id=$1"
	rows, err := database.Con.Query(stmt, id)
	if err != nil {
		common.LogError(err)
	}

	operations, err := mapOperations(rows)
	if err != nil {
		common.LogError(err)
	}

	// Load navigation properties
	operation := operations[0]
	operation.HrdInfo = GetHrd(operation.HrdID.Int64)
	operation.Dispatches = GetDispatches(operation.ID)
	return operation
}

// GetAllOperations returns all operations in the database
func GetAllOperations() []*Operation {
	var stmt = "select * from operations"
	rows, err := database.Con.Query(stmt)

	operations, err := mapOperations(rows)
	if err != nil {
		common.LogError(err)
	}
	return operations
}

// HowManyOperations returnes the total number of operations in the database
func HowManyOperations() int {
	var stmt = "select count(id) as total_operations from operations"
	var totalOperations int

	err := database.Con.QueryRow(stmt).Scan(&totalOperations)

	if err != nil {
		common.LogError(err)
	}

	return totalOperations
}

func mapOperations(rows *sql.Rows) ([]*Operation, error) {
	var err error
	operations := make([]*Operation, 0)

	for rows.Next() {
		o := new(Operation)
		err = rows.Scan(
			&o.ID,
			&o.ProgramID,
			&o.HrdID,
			&o.FscdAnnualPlanID,
			&o.Name,
			&o.Description,
			&o.Year,
			&o.Round,
			&o.Month,
			&o.ExpectedStart,
			&o.ExpectedEnd,
			&o.ActualStart,
			&o.ActualEnd,
			&o.Status,
			&o.CreatedAt,
			&o.UpdatedAt,
			&o.CreatedBy,
			&o.ModifiedBy,
			&o.DeletedAt,
			&o.RationID)

		if err != nil {
			panic(err)
		}
		operations = append(operations, o)
	}

	return operations, err
}
