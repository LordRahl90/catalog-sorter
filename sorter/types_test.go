package sorter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	table := []struct {
		name     string
		args     string
		expected SortOrder
	}{
		{
			name:     "Ascending",
			args:     "asc",
			expected: SortOrderAsc,
		},
		{
			name:     "Descending",
			args:     "desc",
			expected: SortOrderDesc,
		},
		{
			name:     "Unknown",
			args:     "ascd",
			expected: SortOrderUnknown,
		},
		{
			name:     "Unknown",
			args:     "unknown",
			expected: SortOrderUnknown,
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			got := FromString(tt.args)
			assert.Equal(t, tt.expected, got)
		})
	}
}
