package entity

type CartItem struct {
	Id         string  `json:"id"`
	UserId     string  `json:"userId"`
	ProductId  string  `json:"productId"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
}
