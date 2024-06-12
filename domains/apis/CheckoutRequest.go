package apis

type CheckoutItemRequest struct {
	ItemId      string  `json:"item_id"`
	ItemName    string  `json:"item_name"`
	Note        *string `json:"note"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
	VariantId   string  `json:"variant_id"`
	VariantName string  `json:"variant_name"`
}

type CheckoutRequest struct {
	CardNumber  *string `json:"card_number"`
	CustomerId  *string `json:"customer_id"`
	PaymentNote *string `json:"payment_note"`
	PaymentType string  `json:"payment_type"`
	SubTotal    float32 `json:"subtotal"`
	Tendered    float32 `json:"tendered"`

	Items []CheckoutItemRequest `json:"items"`
}
