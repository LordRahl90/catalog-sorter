package sorter

import (
	"fmt"
)

var (
	_ Sorter = (*SorterService)(nil)
)

// SorterService concrete implement of the sorter interface
type SorterService struct {
	plugins    []SorterPlugin
	registered map[string]struct{}
}

// New returns a new sorter implementation
func New() Sorter {
	return &SorterService{
		plugins:    []SorterPlugin{},
		registered: make(map[string]struct{}),
	}
}

// AddPlugin adds a new sorting plugin to the sorter service
// func (ss *SorterService) AddPlugin(plugin SorterPlugin) error {
// 	slugName := slug.Make(plugin.Name())
// 	if _, ok := ss.registered[slugName]; ok {
// 		return fmt.Errorf("plugin already registered: %s", plugin.Name())
// 	}
// 	ss.plugins = append(ss.plugins, plugin)
// 	ss.registered[slugName] = struct{}{}
// 	return nil
// }

// Process processes the provided products by running through the different sorting plugins
func (ss *SorterService) Process(order SortOrder, option SorterPlugin, products ...Product) ([]Product, error) {
	if option == nil {
		return nil, fmt.Errorf("invalid option provided")
	}
	return option.Sort(order, products...)
}
