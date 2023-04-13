package views

import (
	"fmt"
	"sort"

	"github.com/LordRahl90/catalog-sorter/sorter"
)

var (
	_ sorter.SorterPlugin = (*ViewSort)(nil)
)

// ViewSort houses an implementation based on the number of views a product has
type ViewSort struct {
	name    string
	enabled bool
}

// NewViewSort returns a new sorter plugin
func NewViewSort(enabled bool) sorter.SorterPlugin {
	return &ViewSort{
		name:    "View Sort",
		enabled: enabled,
	}
}

// Enabled returns if this plugin is enabled or not
func (v *ViewSort) Enabled() bool {
	return v.enabled
}

// Enable enables the plugin
func (v *ViewSort) Enable() {
	v.enabled = true
}

// Disable flips the switch
func (v *ViewSort) Disable() {
	v.enabled = false
}

// Name returns the name of the of this plugin
func (v *ViewSort) Name() string {
	return v.name
}

// Sort performs the sorting logic based on the provided value
func (v *ViewSort) Sort(order sorter.SortOrder, products ...sorter.Product) ([]sorter.Product, error) {
	if len(products) == 0 {
		return nil, fmt.Errorf("no products to sort")
	}

	switch order {
	case sorter.SortOrderAsc:
		sort.Slice(products, func(i, j int) bool {
			return products[i].ViewCount < products[j].ViewCount
		})
		return products, nil
	case sorter.SortOrderDesc:
		sort.Slice(products, func(i, j int) bool {
			return products[i].ViewCount > products[j].ViewCount
		})
		return products, nil
	default:
		return nil, fmt.Errorf("invalid sorting order provided")
	}
}
