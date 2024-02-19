package entity

type CartItem struct {
	Id     string `json:"id"`
	UserId string `json:"userId"`
	CartItemRequest
}
