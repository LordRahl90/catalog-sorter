package sorter

import (
	"os"
	"sort"
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

func TestNewSortService(t *testing.T) {
	ss := New()

	require.NotNil(t, ss)
}

func TestProcessWithNoPlugin(t *testing.T) {
	products := []Product{
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

	ss := New()
	res, err := ss.Process(SortOrderAsc, nil, products...)
	require.EqualError(t, err, "invalid option provided")
	require.Empty(t, res)
}

func TestProcessWithSinglePlugin(t *testing.T) {
	products := []Product{
		{
			ID:         1,
			SalesCount: 30,
		},
		{
			ID:         2,
			SalesCount: 35,
		},
		{
			ID:         3,
			SalesCount: 25,
		},
	}

	ss := New()

	mp := &pluginMocker{
		NameFunc: func() string {
			return "Mock Plugin"
		},
		EnabledFunc: func() bool {
			return true
		},
	}

	res, err := ss.Process(SortOrderAsc, mp, products...)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, uint64(25), res[0].SalesCount)
	assert.Equal(t, uint64(30), res[1].SalesCount)
	assert.Equal(t, uint64(35), res[2].SalesCount)
}

func TestProcessWithMultiplePlugins(t *testing.T) {
	products := []Product{
		{
			ID:         1,
			SalesCount: 30,
			ViewCount:  40,
		},
		{
			ID:         2,
			SalesCount: 35,
			ViewCount:  38,
		},
		{
			ID:         3,
			SalesCount: 25,
			ViewCount:  36,
		},
	}

	ss := New()

	mp := &pluginMocker{
		NameFunc: func() string {
			return "Mock Plugin"
		},
		EnabledFunc: func() bool {
			return true
		},
	}
	mp2 := &pluginMocker{
		SortFunc: func(order SortOrder, products ...Product) ([]Product, error) {
			sort.Slice(products, func(i, j int) bool {
				return products[i].ViewCount > products[j].ViewCount
			})
			return products, nil
		},
	}

	assert.Equal(t, "Plugin Mocker", mp2.Name())
	assert.True(t, mp2.Enabled())

	res, err := ss.Process(SortOrderAsc, mp, products...)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, uint64(25), res[0].SalesCount)
	assert.Equal(t, uint64(30), res[1].SalesCount)
	assert.Equal(t, uint64(35), res[2].SalesCount)

	res, err = ss.Process(SortOrderUnknown, mp, products...)
	require.EqualError(t, err, "invalid sort order provided")
	require.Nil(t, res)

	res, err = ss.Process(SortOrderAsc, mp2, products...)
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, uint64(36), res[2].ViewCount)
	assert.Equal(t, uint64(38), res[1].ViewCount)
	assert.Equal(t, uint64(40), res[0].ViewCount)
}
