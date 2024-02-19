package entity

type Product struct {
	Id               string   `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"shortDescription,omitempty"`
	Description      string   `json:"description,omitempty"`
	Category         Category `json:"category,omitempty"`
	Price            float64  `json:"price,omitempty"`
	Quantity         int      `json:"quantity,omitempty"`
	ImageUrl         string   `json:"imageUrl"`
}
