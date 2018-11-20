package models

// RegionalRequest represents monthly requests for Relief or PSNP programs
type RegionalRequest struct {
	BaseModel
	ID            int                   `json:"id" bson:"_id"`
	OperationID   int                   `json:"operation_id" bson:"operation_id"`
	ReferenceNo   string                `json:"reference_no" bson:"reference_no"`
	RegionID      int                   `json:"region_id" bson:"region_id"`
	RationID      int                   `json:"ration_id" bson:"ration_id"`
	RequestedDate string                `json:"requested_date" bson:"requested_date"`
	ProgramID     int                   `json:"program_id" bson:"program_id"`
	Remark        string                `json:"remark" bson:"remark"`
	Generated     bool                  `json:"generated" bson:"generated"`
	Items         []RegionalRequestItem `json:"request_items" bson:"request_items"`
}

// RegionalRequestItem represents items within a regional requisition
type RegionalRequestItem struct {
	BaseModel
	ID                    int     `json:"id" bson:"_id"`
	RegionalRequestID     int     `json:"regional_request_id" bson:"regional_request_id"`
	FdpID                 int     `json:"fdp_id" bson:"fdp_id"`
	NumberOfBeneficiaries float64 `json:"number_of_beneficiaries" bson:"number_of_beneficiaries"`
	Fdp                   Fdp     `json:"fdp" bson:"fdp"`
}
