package salesperview

import (
	"os"
	"testing"

	"zabira/sorter"

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
	v := NewSalesPerPriceViewSort(true)
	assert.True(t, v.Enabled())
	assert.Equal(t, "Sales Per View Sort", v.Name())
}

func TestSortingAscendingOrder(t *testing.T) {
	products := []sorter.Product{
		{
			ID:         1,
			SalesCount: 100,
			ViewCount:  20,
		},
		{
			ID:         2,
			SalesCount: 80,
			ViewCount:  80,
		},
		{
			ID:         3,
			SalesCount: 70,
			ViewCount:  35,
		},
	}

	v := NewSalesPerPriceViewSort(true)
	res, err := v.Sort(sorter.SortOrderAsc, products...)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 3)

	assert.Equal(t, res[0].SalesCount, uint64(80))
	assert.Equal(t, res[0].ViewCount, uint64(80))

	assert.Equal(t, res[1].SalesCount, uint64(70))
	assert.Equal(t, res[1].ViewCount, uint64(35))

	assert.Equal(t, res[2].SalesCount, uint64(100))
	assert.Equal(t, res[2].ViewCount, uint64(20))
}

func TestSortingDescendingOrder(t *testing.T) {
	products := []sorter.Product{
		{
			ID:         1,
			SalesCount: 100,
			ViewCount:  20,
		},
		{
			ID:         2,
			SalesCount: 80,
			ViewCount:  80,
		},
		{
			ID:         3,
			SalesCount: 70,
			ViewCount:  35,
		},
	}

	v := NewSalesPerPriceViewSort(true)
	res, err := v.Sort(sorter.SortOrderDesc, products...)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 3)

	assert.Equal(t, res[2].SalesCount, uint64(80))
	assert.Equal(t, res[2].ViewCount, uint64(80))

	assert.Equal(t, res[1].SalesCount, uint64(70))
	assert.Equal(t, res[1].ViewCount, uint64(35))

	assert.Equal(t, res[0].SalesCount, uint64(100))
	assert.Equal(t, res[0].ViewCount, uint64(20))
}

func TestSortEmptySlice(t *testing.T) {
	products := []sorter.Product{}

	v := NewSalesPerPriceViewSort(true)
	res, err := v.Sort(sorter.SortOrderDesc, products...)
	require.EqualError(t, err, "no products to sort")
	assert.Empty(t, res)
}

func TestSortSingleItem(t *testing.T) {
	products := []sorter.Product{
		{
			ID:         1,
			SalesCount: 100,
			ViewCount:  20,
		},
	}

	v := NewSalesPerPriceViewSort(true)
	res, err := v.Sort(sorter.SortOrderDesc, products...)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	assert.Len(t, res, 1)

	assert.Equal(t, res[0].SalesCount, uint64(100))
	assert.Equal(t, res[0].ViewCount, uint64(20))
}

func TestSortItemsWithNoViewCount(t *testing.T) {
	products := []sorter.Product{
		{
			ID:         1,
			SalesCount: 100,
		},
	}

	v := NewSalesPerPriceViewSort(true)
	res, err := v.Sort(sorter.SortOrderDesc, products...)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	assert.Len(t, res, 1)

	assert.Equal(t, res[0].SalesCount, uint64(100))
	assert.Equal(t, res[0].ViewCount, uint64(0))
}

func TestSortItemsWithNoSalesCount(t *testing.T) {
	products := []sorter.Product{
		{
			ID:        1,
			ViewCount: 100,
		},
	}

	v := NewSalesPerPriceViewSort(true)
	res, err := v.Sort(sorter.SortOrderDesc, products...)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	assert.Len(t, res, 1)

	assert.Equal(t, res[0].SalesCount, uint64(0))
	assert.Equal(t, res[0].ViewCount, uint64(100))
}
