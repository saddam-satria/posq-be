package apis

type CreateCustomerRequest struct {
	Address *string `json:"address"`
	Email   string  `json:"email"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
}
