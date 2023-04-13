package price

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
	v := NewPriceSort(true)
	assert.True(t, v.Enabled())
	assert.Equal(t, "Price Sort", v.Name())
}

func TestSortingAscendingOrder(t *testing.T) {
	products := []sorter.Product{
		{
			ID:    1,
			Price: 20.5,
		},
		{
			ID:    2,
			Price: 33.3,
		},
		{
			ID:    3,
			Price: 30.1,
		},
	}

	v := NewPriceSort(true)
	res, err := v.Sort(sorter.SortOrderAsc, products...)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 3)

	assert.Equal(t, res[2].Price, float64(33.3))
	assert.Equal(t, res[1].Price, float64(30.1))
	assert.Equal(t, res[0].Price, float64(20.5))
}

func TestSortingDescendingOrder(t *testing.T) {
	products := []sorter.Product{
		{
			ID:    1,
			Price: 20.5,
		},
		{
			ID:    2,
			Price: 33.3,
		},
		{
			ID:    3,
			Price: 30.1,
		},
	}

	v := NewPriceSort(true)
	res, err := v.Sort(sorter.SortOrderDesc, products...)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Len(t, res, 3)

	assert.Equal(t, res[0].Price, float64(33.3))
	assert.Equal(t, res[1].Price, float64(30.1))
	assert.Equal(t, res[2].Price, float64(20.5))
}

func TestSortEmptySlice(t *testing.T) {
	products := []sorter.Product{}

	v := NewPriceSort(true)
	res, err := v.Sort(sorter.SortOrderDesc, products...)
	require.EqualError(t, err, "no products to sort")
	assert.Empty(t, res)
}

func TestSortSingleItem(t *testing.T) {
	products := []sorter.Product{
		{
			ID:    1,
			Price: 20.5,
		},
	}

	v := NewPriceSort(true)
	res, err := v.Sort(sorter.SortOrderDesc, products...)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	assert.Len(t, res, 1)
	assert.Equal(t, res[0].Price, float64(20.5))
}
