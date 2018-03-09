package models

// RegionalRequest represents monthly requests for Relief or PSNP programs
type RegionalRequest struct {
	BaseModel
	ID            int                   `json:"id"`
	OperationID   int                   `json:"operation_id"`
	ReferenceNo   string                `json:"reference_no"`
	RegionID      int                   `json:"region_id"`
	RationID      int                   `json:"ration_id"`
	RequestedDate string                `json:"requested_date"`
	ProgramID     int                   `json:"program_id"`
	Remark        string                `json:"remark"`
	Generated     bool                  `json:"generated"`
	Items         []RegionalRequestItem `json:"request_items"`
}

// RegionalRequestItem represents items within a regional requisition
type RegionalRequestItem struct {
	BaseModel
	ID                    int     `json:"id"`
	RegionalRequestID     int     `json:"regional_request_id"`
	FdpID                 int     `json:"fdp_id"`
	NumberOfBeneficiaries float64 `json:"number_of_beneficiaries"`
}
