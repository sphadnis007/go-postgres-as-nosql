## About
This simple project implements "Shopping Cart" used on e-commerce websites. To learn more about how it works, checkout the [examples](README.md#examples) section.

## Motive
1. Learn how to interact with Postgres using Golang.
2. Handle NoSQL like data in Postgres using jsonb.
3. Use REST server (gin-gonic) in Golang.

## Pre-requisites
1. Golang should be present on the host.
2. Postgres server must be present/running on the host. It should have configurations as per [configs.go](./configs/configs.go) file (or configs.go should be modified accordingly).

If Docker is present on the host then Postgres container with the required configurations can be started by running the following command:
```bash
docker run -p 5432:5432 -e POSTGRES_PASSWORD=shubham -e POSTGRES_DB=test-db -d postgres
```

## Installation
Clone the repository and simply run the following commands:
```bash
go mod vendor
go run .
```

## API Specifications

List all the carts
```bash
GET /cart
```

List products present in a cart with "cartID"
```bash
GET /cart/:cartID
```

Creates a new cart and adds "numOfObjects" products in it. It generates fake products data and returns cartID to the user. We don't have to provide any data in this call.
```bash
POST /cart/add/:numOfObjects
```

Deletes a product from the cart
```bash
DELETE /cart/:cartID/:productID
```


* Note: All the IDs used in this project are UUIDs.


## Examples

1. Shopping Carts table is empty at the start.
```bash
% curl http://localhost:3000/cart | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100     2  100     2    0     0    250      0 --:--:-- --:--:-- --:--:--   250
[]
```

2. Added 3 carts in the DB with fake data.
```bash
shubham@Shubhams-MacBook-Pro ~ % curl -X POST http://localhost:3000/cart/add/3 | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    51  100    51    0     0   5666      0 --:--:-- --:--:-- --:--:--  5666
"{ cart_id: a79cee22-44a4-447a-87b8-8e5c07214dba }"
shubham@Shubhams-MacBook-Pro ~ % curl -X POST http://localhost:3000/cart/add/1 | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    51  100    51    0     0  10200      0 --:--:-- --:--:-- --:--:-- 10200
"{ cart_id: c3bfa6d3-79be-445f-9e18-bafae2f7b60e }"
shubham@Shubhams-MacBook-Pro ~ % curl -X POST http://localhost:3000/cart/add/2 | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    51  100    51    0     0   7285      0 --:--:-- --:--:-- --:--:--  7285
"{ cart_id: b4b43f97-cdb6-4b34-a173-891356ef6854 }"
```

3. For - curl -X POST http://localhost:3000/cart/add/2 - call, we had received cart id as b4b43f97-cdb6-4b34-a173-891356ef6854. We can check products present in this cart.
```bash
shubham@Shubhams-MacBook-Pro ~ % curl http://localhost:3000/cart/b4b43f97-cdb6-4b34-a173-891356ef6854 | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   324  100   324    0     0  14727      0 --:--:-- --:--:-- --:--:-- 15428
[
    {
        "category": "Objects",
        "created_at": "1900-11-17T05:05:07.501868099Z",
        "id": "7525ee13-92e3-4eda-8e50-550f9bc159d4",
        "name": "Eggplant",
        "price": 7699,
        "quantity": 1
    },
    {
        "category": "People & Body",
        "created_at": "2003-05-16T00:52:15.591212582Z",
        "id": "11531d65-0759-460a-9941-07407d223d9b",
        "name": "Carrot",
        "price": 1469,
        "quantity": 1
    }
]
```

4. Deleting "Carrot" from the above cart
```bash
shubham@Shubhams-MacBook-Pro ~ % curl -X DELETE http://localhost:3000/cart/b4b43f97-cdb6-4b34-a173-891356ef6854/11531d65-0759-460a-9941-07407d223d9b | python -m json.tool 
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    34  100    34    0     0   1307      0 --:--:-- --:--:-- --:--:--  1307
"Product Removed from the Cart!!!"
```

5. Rechecking if "Carrot" got removed or not.
```bash
shubham@Shubhams-MacBook-Pro ~ % curl http://localhost:3000/cart/b4b43f97-cdb6-4b34-a173-891356ef6854 | python -m json.tool
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   158  100   158    0     0  39500      0 --:--:-- --:--:-- --:--:-- 39500
[
    {
        "category": "Objects",
        "created_at": "1900-11-17T05:05:07.501868099Z",
        "id": "7525ee13-92e3-4eda-8e50-550f9bc159d4",
        "name": "Eggplant",
        "price": 7699,
        "quantity": 1
    }
]
```

## Non-standard Modules used
1. Postgres driver - [pgx](https://github.com/jackc/pgx)
2. REST server - [gin-gonic](https://github.com/gin-gonic/gin)
3. Fake data generator - [gofakeit](https://github.com/brianvoe/gofakeit)
4. UUID - [google/uuid](https://github.com/google/uuid)
