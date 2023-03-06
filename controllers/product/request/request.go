package request

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Type        string  `json:"type"`
}
