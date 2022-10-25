
# Project Description

## Introduction
The primary focus for the code are: clean, simple, readable, maintainable, and testable. It should be easy to glance the code, then easily and quickly understand what it is about.

## The Task

The mini project is to implement a simple warehouse REST API service. This API works on products management, and each product contains an identification number, a name, a price and an available stock. It is to load products into the software from a file: products.json.

The warehouse has the following functionality;

* Get all products with a formatted price.
* Get a specific product with formatted price.
* Update the product stock.
* Create a product.


## set up project

### docker
run command under root folder: `docker-compose up` 
### data migration
run command under root folder: `docker run -v /Users/taosun/Documents/continue-forward/go/rest-api-product-mangagement/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://root:password@localhost:5432/rest_api?sslmode=disable" up 1`
