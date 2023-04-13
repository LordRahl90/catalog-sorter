package sorter

import (
	"fmt"
	"sort"
)

var (
	_ SorterPlugin = (*pluginMocker)(nil)
)

type pluginMocker struct {
	EnabledFunc func() bool
	NameFunc    func() string
	SortFunc    func(order SortOrder, products ...Product) ([]Product, error)
}

// Enabled implements SorterPlugin
func (pm *pluginMocker) Enabled() bool {
	if pm.EnabledFunc == nil {
		return true // always enabled
	}
	return pm.EnabledFunc()
}

// Name implements SorterPlugin
func (pm *pluginMocker) Name() string {
	if pm.NameFunc == nil {
		return "Plugin Mocker"
	}
	return pm.NameFunc()
}

// Sort implements SorterPlugin
func (pm *pluginMocker) Sort(order SortOrder, products ...Product) ([]Product, error) {
	if pm.SortFunc == nil {
		return Sort(order, products...)
	}
	return pm.SortFunc(order, products...)
}

func Sort(order SortOrder, products ...Product) ([]Product, error) {
	switch order {
	case SortOrderAsc:
		sort.Slice(products, func(i, j int) bool {
			return products[i].SalesCount < products[j].SalesCount
		})
		return products, nil
	case SortOrderDesc:
		sort.Slice(products, func(i, j int) bool {
			return products[i].SalesCount > products[j].SalesCount
		})
		return products, nil

	default:
		return nil, fmt.Errorf("invalid sort order provided")
	}
}
