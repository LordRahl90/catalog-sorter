package sorter

// Product entity structs for holding the product
type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Created    string  `json:"created"`
	SalesCount uint64  `json:"sales_count"`
	ViewCount  uint64  `json:"views_count"`
}
