package adapter

import (
	"github.com/LordRahl90/catalog-sorter/entity"
	"github.com/LordRahl90/catalog-sorter/sorter"
)

// Product adapters for handling product comparisons
type SPVProduct struct {
	Product      entity.Product
	SalesPerView float64
}

// Products array of product pointers
type Products []*SPVProduct

func (p Products) Len() int {
	return len(p)
}

// Less determine if the values in the indices are lesser
func (p Products) Less(i, j int) bool {
	return p[i].SalesPerView < p[j].SalesPerView
}

// Swap swaps the values of the 2 indices
func (p Products) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// FromEntity converts products from sorter entity to package level product
func FromEntity(products []sorter.Product) Products {
	result := make(Products, len(products))
	for i, v := range products {
		result[i] = &SPVProduct{
			Product: entity.Product{
				ID:         v.ID,
				Name:       v.Name,
				Price:      v.Price,
				Created:    v.Created,
				SalesCount: v.SalesCount,
				ViewCount:  v.ViewCount,
			},
		}

		if v.SalesCount > 0 && v.ViewCount > 0 {
			result[i].SalesPerView = float64(v.SalesCount) / float64(v.ViewCount)
		}
	}

	return result
}

// ToEntity converts this domain product to sorter product
func ToEntity(products Products) []sorter.Product {
	result := make([]sorter.Product, products.Len())
	for i, v := range products {
		result[i] = sorter.Product{
			ID:         v.Product.ID,
			Name:       v.Product.Name,
			Price:      v.Product.Price,
			Created:    v.Product.Created,
			SalesCount: v.Product.SalesCount,
			ViewCount:  v.Product.ViewCount,
		}
	}
	return result
}
