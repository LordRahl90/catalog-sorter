package views

import (
	"os"
	"testing"

	"github.com/LordRahl90/catalog-sorter/sorter"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	code := 1
	defer func() {
		os.Exit(code)
	}()

	code = m.Run()
}

func TestCreateViewName(t *testing.T) {
	v := NewViewSort(true)
	assert.True(t, v.Enabled())
	assert.Equal(t, "View Sort", v.Name())
}

func TestSortingAscendingOrder(t *testing.T) {
	products := []sorter.Product{
		{
			ID:        1,
			ViewCount: 20,
		},
		{
			ID:        2,
			ViewCount: 30,
		},
		{
			ID:        3,
			ViewCount: 25,
		},
	}

	v := NewViewSort(true)
	res, err := v.Sort(sorter.SortOrderAsc, products...)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 3)

	assert.Equal(t, res[0].ViewCount, uint64(20))
	assert.Equal(t, res[1].ViewCount, uint64(25))
	assert.Equal(t, res[2].ViewCount, uint64(30))
}

func TestSortingDescendingOrder(t *testing.T) {
	products := []sorter.Product{
		{
			ID:        1,
			ViewCount: 20,
		},
		{
			ID:        2,
			ViewCount: 30,
		},
		{
			ID:        3,
			ViewCount: 25,
		},
	}

	v := NewViewSort(true)
	res, err := v.Sort(sorter.SortOrderDesc, products...)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 3)

	assert.Equal(t, res[0].ViewCount, uint64(30))
	assert.Equal(t, res[1].ViewCount, uint64(25))
	assert.Equal(t, res[2].ViewCount, uint64(20))
}

func TestSortEmptySlice(t *testing.T) {
	products := []sorter.Product{}

	v := NewViewSort(true)
	res, err := v.Sort(sorter.SortOrderDesc, products...)
	require.EqualError(t, err, "no products to sort")
	assert.Empty(t, res)
}
