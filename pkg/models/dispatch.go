package models

import (
	"github.com/ndrmc/rm2doc/pkg/common"
	"github.com/ndrmc/rm2doc/pkg/database"
)

// Dispatch represent a single goods issue ticket from warehouse
type Dispatch struct {
	BaseModel
	ID                   int            `json:"id"`
	GinNo                string         `json:"gin_no"`
	OperationID          int            `json:"operation_id"`
	RequisitionNo        string         `json:"requisition_no"`
	DispatchDate         string         `json:"dispatch_date"`
	FdpID                int            `json:"fdpi_d"`
	WeightBridgeTicketNo string         `json:"weight_bridge_ticket_no"`
	TransporterID        int            `json:"transporter_id"`
	PlateNo              string         `json:"plate_no"`
	TrailerPlateNo       string         `json:"trailer_plate_no"`
	DriverName           string         `json:"driver_name"`
	Remark               string         `json:"remark"`
	Draft                bool           `json:"draft"`
	Deleted              bool           `json:"deleted"`
	HubID                int            `json:"hub_id"`
	WarehouseID          int            `json:"warehouse_id"`
	StorekeeperName      string         `json:"storekeeper_name"`
	DispatchIDGUID       string         `json:"dispatch_id_guid"`
	DispatchedDateEC     string         `json:"dispatched_date_ec"`
	DispatchTypeID       int            `json:"dispatch_type_id"`
	DispatchType         int            `json:"dispatch_type"`
	Items                []DispatchItem `json:"dispatch_items"`
}

//DispatchItem represents item detail for a dispatch
type DispatchItem struct {
	BaseModel
	ID                  int     `json:"id"`
	DispatchID          int     `json:"dispatch_id"`
	CommodityCategoryID int     `json:"commodity_category_id"`
	CommodityID         int     `json:"commodity_id"`
	Quantity            float32 `json:"quantity"`
	ProjectID           int     `json:"project_id"`
	GUIDRef             string  `json:"guid_ref"`
	OrganizationID      int     `json:"organization_id"`
	UnitOfMeasureID     int     `json:"unit_of_measure_id"`
	Deleted             bool    `json:"deleted"`
}

// GetDispatch fetches a specific Dispatch record from transactional database
func GetDispatch(id int) Dispatch {
	var dispatch Dispatch
	database.Session.Find(&dispatch, id)
	return dispatch
}

// GetDispatches fetches all dispatch records for a given operation
func GetDispatches(operationID int) []Dispatch {
	var dispatches []Dispatch
	database.Session.Where("operation_id=?", operationID).Find(dispatches)
	return dispatches
}

// GetAllDispatches returns all dispatches
func GetAllDispatches() []Dispatch {
	var dispatches []Dispatch
	database.Session.Preload("Items").Find(&dispatches)
	return dispatches
}

// TotalDispatch calculates total dispatch amount
func TotalDispatch(operationID int) float32 {
	var total float32
	errs := database.Session.Raw(`SELECT sum(di.quantity) as quantity_sum
						  FROM dispatch_items dd
						  INNER JOIN dispatches d ON di.dispatch_id = d.id
						  WHERE d.operation_id = ?`, operationID).Scan(&total).GetErrors()

	for _, err := range errs {
		common.LogError(err)
	}

	return total
}
