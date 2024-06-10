package model

type Item struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

type Discount struct {
	Type         string  `json:"type"`
	Amount       float64 `json:"amount"`
	Percentage   float64 `json:"percentage"`
	Category     string  `json:"category"`
	Points       int     `json:"points"`
	EveryXTHB    float64 `json:"every_x_thb"`
	DiscountYTHB float64 `json:"discount_y_thb"`
}

type Cart struct {
	Items     []Item     `json:"items"`
	Discounts []Discount `json:"discounts"`
}
