package main

import (
	"os"
	"testing"

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

func TestInput(t *testing.T) {
	path := "./testdata/products.json"
	res, err := input(path)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Len(t, res, 3)
}

func TestHandleHappyPath(t *testing.T) {
	path := "./testdata/products.json"
	products, err := input(path)
	require.NoError(t, err)
	require.NotNil(t, products)

	assert.Len(t, products, 3)

	res, err := handle("asc", "price-sorter", products)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Len(t, res, 3)

	assert.Equal(t, float64(10.00), res[0].Price)
	assert.Equal(t, float64(12.99), res[1].Price)
	assert.Equal(t, float64(44.49), res[2].Price)
}

func TestHandleHappyPathDesc(t *testing.T) {
	path := "./testdata/products.json"
	products, err := input(path)
	require.NoError(t, err)
	require.NotNil(t, products)

	assert.Len(t, products, 3)

	res, err := handle("desc", "price-sorter", products)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Len(t, res, 3)

	assert.Equal(t, float64(10.00), res[2].Price)
	assert.Equal(t, float64(12.99), res[1].Price)
	assert.Equal(t, float64(44.49), res[0].Price)
}

func TestHandleInvalidSorter(t *testing.T) {
	path := "./testdata/products.json"
	products, err := input(path)
	require.NoError(t, err)
	require.NotNil(t, products)

	assert.Len(t, products, 3)

	res, err := handle("asc", "price-sorters", products)
	require.EqualError(t, err, "invalid plugin provided")
	require.Nil(t, res)
}

func TestHandleInvalidSortOrder(t *testing.T) {
	path := "./testdata/products.json"
	products, err := input(path)
	require.NoError(t, err)
	require.NotNil(t, products)

	assert.Len(t, products, 3)

	res, err := handle("ascx", "price-sorter", products)
	require.EqualError(t, err, "invalid sorting order provided")
	require.Nil(t, res)
}
