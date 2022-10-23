
# Project Description

## Introduction
The primary focus for the code are: simplicity, readability, maintainability, testability. It should be easy to scan the code, and quickly understand what it’s doing.

## The Task

The mini project is to implement a simple warehouse REST API. This API work on products, and the product contains an identification number, a name, a price and an available stock. It is to load products into the software from a file: products.json.

The warehouse has the following functionality;

* Get all products with a formatted price.
* Get a specific product with formatted price.
* Update the product stock.
* Create a product.


## set up project

### docker
run command: `docker-compose`
### data migration
run command under root: `docker run -v /Users/taosun/Documents/continue-forward/go/rest-api-product-mangagement/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://root:password@localhost:5432/rest_api?sslmode=disable" up 1`
