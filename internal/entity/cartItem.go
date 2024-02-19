package entity

type CartItem struct {
	Id         string  `json:"id"`
	UserId     string  `json:"userId,omitempty"`
	Product    Product `json:"product"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
}
