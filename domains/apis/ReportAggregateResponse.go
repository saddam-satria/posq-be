package apis

type ReportAggregateResponse struct {
	ProductOutStock    int `json:"product_out_of_stock"`
	TotalOrder         int `json:"total_order_today"`
	UnavailableService int `json:"unavailable_service"`
}
