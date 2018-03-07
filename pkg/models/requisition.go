package models

import (
	"database/sql"
	"time"

	"github.com/ndrmc/analytics/pkg/common"
	"github.com/ndrmc/analytics/pkg/database"
)

// Requisition represents allocation request document
type Requisition struct {
	Base
	ID            int                `json:"id"`
	RequisitionNo string             `json:"requisition_no"`
	OperationID   sql.NullInt64      `json:"operation_id"`
	CommodityID   sql.NullInt64      `json:"commodity_id"`
	RegionID      sql.NullInt64      `json:"region_id"`
	ZoneID        sql.NullInt64      `json:"zone_id"`
	RationID      sql.NullInt64      `json:"ration_id"`
	RequestedBy   sql.NullString     `json:"requested_by"`
	RequestedOn   time.Time          `json:"requested_on"`
	Status        sql.NullInt64      `json:"status"`
	RequestID     sql.NullString     `json:"request_id"`
	Items         []*RequisitionItem `json:"requisition_items"`
}

// RequisitionItem represents beneficiary number for each FDP together with commodity
type RequisitionItem struct {
	Base
	ID            int             `json:"id"`
	RequisitionID sql.NullInt64   `json:"requisition_id"`
	FdpID         sql.NullInt64   `json:"fdp_id"`
	BeneficiaryNo int             `json:"beneficiary_no"`
	Amount        sql.NullFloat64 `json:"amount"`
	Contigency    sql.NullFloat64 `json:"contigency"`
}

// GetRequisition returns the complete detail for a requisition record
func GetRequisition(id int64) *Requisition {
	var stmt = "select * from requisitions where id=$1"
	rows, err := database.Con.Query(stmt, id)
	if err != nil {
		common.LogError(err)
	}

	requisitions, err := mapRequisitions(rows)
	if err != nil {
		common.LogError(err)
	}

	return requisitions[0]
}

// GetRequisitions fetches all requisitions records for a given operation
func GetRequisitions(operationID int) []*Requisition {
	var stmt = "select * from requisitions where operation_id=$1"
	rows, err := database.Con.Query(stmt, operationID)
	if err != nil {
		common.LogError(err)
	}

	requisitions, err := mapRequisitions(rows)
	if err != nil {
		common.LogError(err)
	}

	return requisitions
}

// TotalBeneficiaries returns the total number of beneficiaries in an operation
func TotalBeneficiaries(operationID int) float64 {
	var stmt = `SELECT sum(ri.beneficiary_no) AS total_beneficiaries
				FROM requisition_items ri
				INNER JOIN requisitions r ON ri.requisition_id = r.id
				WHERE r.operation_id = $1`
	var beneficiaries float64
	err := database.Con.QueryRow(stmt, operationID).Scan(&beneficiaries)

	if err != nil {
		common.LogError(err)
	}
	return beneficiaries
}

func getRequisitionItems(requisitionID int) []*RequisitionItem {
	var stmt = "select * from requisition_items where requisition_id=$1"
	rows, err := database.Con.Query(stmt, requisitionID)
	if err != nil {
		common.LogError(err)
	}

	items, err := mapRequisitionItems(rows)
	if err != nil {
		common.LogError(err)
	}
	return items
}

func mapRequisitions(rows *sql.Rows) ([]*Requisition, error) {
	var err error
	requisitions := make([]*Requisition, 0)

	for rows.Next() {
		r := new(Requisition)
		err = rows.Scan(
			&r.ID,
			&r.RequisitionNo,
			&r.OperationID,
			&r.CommodityID,
			&r.RegionID,
			&r.ZoneID,
			&r.RationID,
			&r.RequestedBy,
			&r.RequestedOn,
			&r.Status,
			&r.CreatedAt,
			&r.UpdatedAt,
			&r.CreatedBy,
			&r.ModifiedBy,
			&r.DeletedAt,
			&r.RequestID)

		if err != nil {
			panic(err)
		}

		r.Items = getRequisitionItems(r.ID)
		requisitions = append(requisitions, r)
	}

	return requisitions, err
}

func mapRequisitionItems(rows *sql.Rows) ([]*RequisitionItem, error) {
	var err error
	items := make([]*RequisitionItem, 0)

	for rows.Next() {
		ri := new(RequisitionItem)
		err = rows.Scan(
			&ri.ID,
			&ri.RequisitionID,
			&ri.FdpID,
			&ri.BeneficiaryNo,
			&ri.Amount,
			&ri.Contigency,
			&ri.CreatedAt,
			&ri.UpdatedAt,
			&ri.CreatedBy,
			&ri.ModifiedBy,
			&ri.DeletedAt)

		if err != nil {
			panic(err)
		}
		// r.Items = getDispatchItems(d.ID)
		items = append(items, ri)
	}

	return items, err
}
