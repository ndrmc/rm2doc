package models

import (
	"github.com/ndrmc/rm2doc/pkg/database"
)

// Requisition represents allocation request document
type Requisition struct {
	BaseModel
	ID            int               `json:"id" bson:"_id"`
	RequisitionNo string            `json:"requisition_no" bson:"requisition_no"`
	OperationID   int               `json:"operation_id" bson:"operation_id"`
	CommodityID   int               `json:"commodity_id" bson:"commodity_id"`
	RegionID      int               `json:"region_id" bson:"region_id"`
	ZoneID        int               `json:"zone_id" bson:"zone_id"`
	RationID      int               `json:"ration_id" bson:"ration_id"`
	RequestedBy   string            `json:"requested_by" bson:"requested_by"`
	RequestedOn   string            `json:"requested_on" bson:"requested_on"`
	Status        int               `json:"status" bson:"status"`
	RequestID     string            `json:"request_id" bson:"request_id"`
	Items         []RequisitionItem `json:"requisition_items" bson:"requisition_items"`
}

// RequisitionItem represents beneficiary number for each FDP together with commodity
type RequisitionItem struct {
	BaseModel
	ID            int     `json:"id" bson:"_id"`
	RequisitionID int     `json:"requisition_id" bson:"requisition_id"`
	FdpID         int     `json:"fdp_id" bson:"fdp_id"`
	BeneficiaryNo int     `json:"beneficiary_no" bson:"beneficiary_no"`
	Amount        float32 `json:"amount" bson:"amount"`
	Contigency    float32 `json:"contigency" bson:"contigency"`
	Fdp           Fdp     `json:"fdp" bson:"fdp"`
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
