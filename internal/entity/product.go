package entity

type Product struct {
	Id               string   `json:"id" param:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"shortDescription"`
	Description      string   `json:"description"`
	Category         Category `json:"category"`
	Price            float64  `json:"price"`
	Quantity         int      `json:"quantity"`
	ImageUrl         string   `json:"imageUrl"`
}
