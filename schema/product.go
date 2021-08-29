package schema

import "time"

type Product struct {
	ID        string    `json:"id" fake:"{uuid}"`
	Name      string    `json:"name" fake:"{vegetable}"`
	Category  string    `json:"category" fake:"{emojicategory}"`
	Price     int       `json:"price" fake:"{number:100,10000}"`
	Quantity  int       `json:"quantity" fake:"{number:1,5}"`
	CreatedAt time.Time `json:"created_at"`
}
