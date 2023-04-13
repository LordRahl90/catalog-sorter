## Extensible Sorting 

### Setup

* Clone repository
* run `go mod tidy`
* run `make run order=asc plugin=svp path="cmd/testdata/products.json"` <br />
NB: The variables can be alternated to various options

Using the provided json data, the following result should be expected:

```json
[
        {
                "id": 1,
                "name": "Alabaster Table",
                "price": 12.99,
                "created": "2019-01-04",
                "sales_count": 32,
                "views_count": 730
        },
        {
                "id": 3,
                "name": "Coffee Table",
                "price": 10,
                "created": "2014-05-28",
                "sales_count": 1048,
                "views_count": 20123
        },
        {
                "id": 2,
                "name": "Zebra Table",
                "price": 44.49,
                "created": "2012-01-04",
                "sales_count": 301,
                "views_count": 3279
        }
]
```


### Contribution

To add a new plugin, create a package that implements the `sorter.SortPlugin` interface, then register it in the `plugins` map in `main.go` using a unique name as the key and the new plugin service as the value.

