package apis

type CheckoutResponse struct {
	Changed     float32 `json:"changed"`
	ReferenceId string  `json:"reference_id"`
	Tendered    float32 `json:"tendered"`
	Total       float32 `json:"total"`
}
