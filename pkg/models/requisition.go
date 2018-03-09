package models

import (
	"github.com/ndrmc/analytics/pkg/database"
)

// Requisition represents allocation request document
type Requisition struct {
	BaseModel
	ID            int               `json:"id"`
	RequisitionNo string            `json:"requisition_no"`
	OperationID   int               `json:"operation_id"`
	CommodityID   int               `json:"commodity_id"`
	RegionID      int               `json:"region_id"`
	ZoneID        int               `json:"zone_id"`
	RationID      int               `json:"ration_id"`
	RequestedBy   string            `json:"requested_by"`
	RequestedOn   string            `json:"requested_on"`
	Status        int               `json:"status"`
	RequestID     string            `json:"request_id"`
	Items         []RequisitionItem `json:"requisition_items"`
}

// RequisitionItem represents beneficiary number for each FDP together with commodity
type RequisitionItem struct {
	BaseModel
	ID            int     `json:"id"`
	RequisitionID int     `json:"requisition_id"`
	FdpID         int     `json:"fdp_id"`
	BeneficiaryNo int     `json:"beneficiary_no"`
	Amount        float32 `json:"amount"`
	Contigency    float32 `json:"contigency"`
	Fdp           Fdp     `json:"fdp"`
}

// GetRequisition returns the complete detail for a requisition record
func GetRequisition(id int) Requisition {
	var req Requisition
	database.Session.Find(&req, id)

	return req
}

// GetRequisitions fetches all requisitions records for a given operation
func GetRequisitions(operationID int) []Requisition {
	var reqs []Requisition
	database.Session.Where("operation_id=?", operationID).Find(&reqs)
	return reqs
}

// TotalBeneficiaries returns the total number of beneficiaries in an operation
func TotalBeneficiaries(operationID int) float32 {
	type SUM struct {
		TotalBeneficiaries float32
	}

	var sum SUM
	database.Session.Raw(`SELECT sum(ri.beneficiary_no) AS total_beneficiaries
				FROM requisition_items ri
				INNER JOIN requisitions r ON ri.requisition_id = r.id
				WHERE r.operation_id = ?`, operationID).Scan(&sum)
	return sum.TotalBeneficiaries
}
