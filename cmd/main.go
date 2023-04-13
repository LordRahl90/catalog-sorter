package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/LordRahl90/catalog-sorter/plugins/price"
	"github.com/LordRahl90/catalog-sorter/plugins/salesperview"
	"github.com/LordRahl90/catalog-sorter/plugins/views"
	"github.com/LordRahl90/catalog-sorter/sorter"
)

var (
	pathStr *string
	order   *string
	plugin  *string

	plugins = map[string]sorter.SorterPlugin{
		"view-sorter":  views.NewViewSort(true),
		"price-sorter": price.NewPriceSort(true),
		"svp":          salesperview.NewSalesPerPriceViewSort(true),
	}

	handler = sorter.New()
)

func main() {
	pathStr = flag.String("path", "./cmd/testdata/products.json", "The path to the products json file")
	order = flag.String("order", "asc", "Order in which the values should be sorted")
	plugin = flag.String("plugin", "view-sorter", "The short name of the desired sorting plugin to use")

	flag.Parse()

	products, err := input(*pathStr)
	if err != nil {
		log.Fatal(err)
	}

	res, err := handle(*order, *plugin, products)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(res, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n", b)
}

func handle(orderStr string, plugin string, products []sorter.Product) ([]sorter.Product, error) {
	v, ok := plugins[plugin]
	if !ok {
		return nil, fmt.Errorf("invalid plugin provided")
	}
	order := sorter.FromString(orderStr)
	if order == sorter.SortOrderUnknown {
		return nil, fmt.Errorf("invalid sorting order provided")
	}

	return handler.Process(order, v, products...)
}

// processes a json data and returns an array products
func input(path string) ([]sorter.Product, error) {
	var result []sorter.Product
	// TODO: Bad idea loading the entire file into memory
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(f, &result); err != nil {
		return nil, err
	}
	return result, nil
}
