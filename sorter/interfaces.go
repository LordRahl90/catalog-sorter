package sorter

// Sorter defines a sorter interface which allows the implementing type to provide concrete sorting implementation
type SorterPlugin interface {
	Name() string
	Enabled() bool
	Sort(order SortOrder, products ...Product) ([]Product, error)
}

// Sorter interface for the sorting service.
type Sorter interface {
	// AddPlugin(plugin SorterPlugin) error
	Process(order SortOrder, option SorterPlugin, producs ...Product) ([]Product, error)
}
