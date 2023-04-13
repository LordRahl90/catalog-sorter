package entity

// Product entity describing the basic product definition
type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    string
	SalesCount uint64
	ViewCount  uint64
}
