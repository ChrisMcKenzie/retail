# Retail PoC

## Requirements

-   Golang `>1.19`

## Running

To get started simply run the following

```
go run main.go
```

The server will be listening by default on `localhost:8080`

## Adding prices

to add a price to a given product id you can run the following.

```
./scripts/add_price.sh <id> <price> <currency>
```

the options available for currency ar USD, EUR, GBP, CAD
