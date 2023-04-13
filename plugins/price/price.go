package price

import (
	"fmt"
	"sort"

	"github.com/LordRahl90/catalog-sorter/plugins/price/adapter"
	"github.com/LordRahl90/catalog-sorter/sorter"
)

var (
	_ sorter.SorterPlugin = (*PriceSort)(nil)
)

// PriceSort houses a sorting implementation based on the price of a product
type PriceSort struct {
	name    string
	enabled bool
}

// NewViewSort returns a new sorter plugin
func NewPriceSort(enabled bool) sorter.SorterPlugin {
	return &PriceSort{
		name:    "Price Sort",
		enabled: enabled,
	}
}

// Enabled returns if this plugin is enabled or not
func (v *PriceSort) Enabled() bool {
	return v.enabled
}

// Enable enables the plugin
func (v *PriceSort) Enable() {
	v.enabled = true
}

// Disable flips the switch
func (v *PriceSort) Disable() {
	v.enabled = false
}

// Name returns the name of the of this plugin
func (v *PriceSort) Name() string {
	return v.name
}

// Sort performs the sorting logic based on the provided value
func (v *PriceSort) Sort(order sorter.SortOrder, products ...sorter.Product) ([]sorter.Product, error) {
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
