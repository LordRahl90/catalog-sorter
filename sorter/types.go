package sorter

type SortOrder int

const (
	// SortOrderUnknown unknown sort order
	SortOrderUnknown SortOrder = iota
	// SortOrderAsc ascending order setting indicator
	SortOrderAsc
	// SortOrderDesc descending order setting indicator
	SortOrderDesc
)

// FromString returns a sort order from a string representation
func FromString(s string) SortOrder {
	switch s {
	case "asc":
		return SortOrderAsc
	case "desc":
		return SortOrderDesc
	default:
		return SortOrderUnknown
	}
}
