package entity

import "log"

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsAvailable bool    `json:"isAvailable"`
}

func main() {
	shoe := Product{
		ID:          `1002`,
		Name:        "Nike",
		Description: "A shoe for athletes",
		Price:       199.99,
		IsAvailable: true,
	}

	log.Printf("shoe: %+v\n", shoe)
}
