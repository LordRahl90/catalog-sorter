package adapter

import (
	"zabira/sorter"
)

// Product adapters for handling product comparisons
type Product struct {
	ID         int
	Name       string
	Price      float64
	Created    string
	SalesCount uint64
	ViewCount  uint64
}

// Products array of product pointers
type Products []*Product

func (p Products) Len() int {
	return len(p)
}

func (p Products) Less(i, j int) bool {
	return p[i].Price < p[j].Price
}

func (p Products) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// FromEntity converts products from sorter entity to package level product
func FromEntity(products []sorter.Product) Products {
	result := make(Products, len(products))
	for i, v := range products {
		result[i] = &Product{
			ID:         v.ID,
			Name:       v.Name,
			Price:      v.Price,
			Created:    v.Created,
			SalesCount: v.SalesCount,
			ViewCount:  v.ViewCount,
		}
	}

	return result
}

// ToEntity converts this domain product to sorter product
func ToEntity(products Products) []sorter.Product {
	result := make([]sorter.Product, products.Len())
	for i, v := range products {
		result[i] = sorter.Product{
			ID:         v.ID,
			Name:       v.Name,
			Price:      v.Price,
			Created:    v.Created,
			SalesCount: v.SalesCount,
			ViewCount:  v.ViewCount,
		}
	}
	return result
}
