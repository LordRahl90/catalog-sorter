package salesperview

import (
	"fmt"
	"sort"

	"github.com/LordRahl90/catalog-sorter/plugins/salesperview/adapter"
	"github.com/LordRahl90/catalog-sorter/sorter"
)

// SalesPriceViewSort houses a sorting implementation based on the price of a product
type SalesPerPriceViewSort struct {
	name    string
	enabled bool
}

// NewSalesPerPriceViewSort returns a new sorter plugin
func NewSalesPerPriceViewSort(enabled bool) sorter.SorterPlugin {
	return &SalesPerPriceViewSort{
		name:    "Sales Per View Sort",
		enabled: enabled,
	}
}

// Enabled returns if this plugin is enabled or not
func (spv *SalesPerPriceViewSort) Enabled() bool {
	return spv.enabled
}

// Name returns the name of the of this plugin
func (spv *SalesPerPriceViewSort) Name() string {
	return spv.name
}

// Sort performs the sorting logic based on the provided value
func (spv *SalesPerPriceViewSort) Sort(order sorter.SortOrder, products ...sorter.Product) ([]sorter.Product, error) {
	if len(products) == 0 {
		return nil, fmt.Errorf("no products to sort")
	}
	adapterProduct := adapter.FromEntity(products)
	sort.Sort(adapterProduct)

	switch order {
	case sorter.SortOrderAsc:
		return adapter.ToEntity(adapterProduct), nil
	case sorter.SortOrderDesc:
		sort.Sort(sort.Reverse(adapterProduct))
		return adapter.ToEntity(adapterProduct), nil
	default:
		return nil, fmt.Errorf("invalid sorting order provided")
	}
}
