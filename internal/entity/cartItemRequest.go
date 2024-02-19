package entity

type CartItemRequest struct {
	ProductId  string  `json:"productId"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
}
