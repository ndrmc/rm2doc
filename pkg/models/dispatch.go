package models

import (
	"database/sql"
	"time"

	"github.com/ndrmc/analytics/pkg/common"
	"github.com/ndrmc/analytics/pkg/database"
)

// Dispatch represent a single goods issue ticket from warehouse
type Dispatch struct {
	Base
	ID                   int            `json:"id"`
	GinNo                int            `json:"gin_no"`
	OperationID          sql.NullInt64  `json:"operation_id"`
	RequisitionNo        sql.NullString `json:"requisition_no"`
	DispatchDate         time.Time      `json:"dispatch_date"`
	FdpID                sql.NullInt64  `json:"fdp_id"`
	WeightBridgeTicketNo sql.NullString `json:"weight_bridge_ticket_no"`
	TransporterID        sql.NullInt64  `json:"transporter_id"`
	PlateNo              sql.NullString `json:"plate_no"`
	TrailerPlateNo       sql.NullString `json:"trailer_plate_no"`
	DriverName           sql.NullString `json:"driver_name"`
	Remark               sql.NullString `json:"remark"`
	Draft                sql.NullBool   `json:"draft"`
	Deleted              sql.NullBool   `json:"deleted"`
	HubID                sql.NullInt64  `json:"hub_id"`
	WarehouseID          sql.NullInt64  `json:"warehouse_id"`
	StorekeeperName      sql.NullString `json:"storekeeper_name"`
	DispatchIDGUID       sql.NullString `json:"dispatch_id_guid"`
	DispatchedDateEC     sql.NullString `json:"dispatched_date_ec"`
	DispatchTypeID       sql.NullInt64  `json:"dispatch_type_id"`
	DispatchType         sql.NullInt64  `json:"dispatch_type"`
}

// GetDispatch fetches a specific Dispatch record from transactional database
func GetDispatch(id int) *Dispatch {
	var stmt = "select * from dispatches where id=$1"
	rows, err := database.Con.Query(stmt, id)
	if err != nil {
		common.LogError(err)
	}

	dispatches, err := mapDispatches(rows)
	if err != nil {
		common.LogError(err)
	}

	return dispatches[0]
}

// GetDispatches fetches all dispatch records for a given operation
func GetDispatches(operationID int64) []*Dispatch {
	var stmt = "select * from dispatches where operation_id=$1"
	rows, err := database.Con.Query(stmt, operationID)
	if err != nil {
		common.LogError(err)
	}

	dispatches, err := mapDispatches(rows)
	if err != nil {
		common.LogError(err)
	}

	return dispatches
}

func mapDispatches(rows *sql.Rows) ([]*Dispatch, error) {
	var err error
	dispatches := make([]*Dispatch, 0)

	for rows.Next() {
		d := new(Dispatch)
		err = rows.Scan(
			&d.ID,
			&d.GinNo,
			&d.OperationID,
			&d.RequisitionNo,
			&d.DispatchDate,
			&d.FdpID,
			&d.WeightBridgeTicketNo,
			&d.TransporterID,
			&d.PlateNo,
			&d.TrailerPlateNo,
			&d.DriverName,
			&d.Remark,
			&d.Draft,
			&d.CreatedBy,
			&d.ModifiedBy,
			&d.Deleted,
			&d.DeletedAt,
			&d.CreatedAt,
			&d.UpdatedAt,
			&d.HubID,
			&d.WarehouseID,
			&d.StorekeeperName,
			&d.DispatchIDGUID,
			&d.DispatchedDateEC,
			&d.DispatchTypeID,
			&d.DispatchType)

		if err != nil {
			panic(err)
		}
		dispatches = append(dispatches, d)
	}

	return dispatches, err
}
