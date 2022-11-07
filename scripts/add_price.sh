#!/bin/bash

id=$1
price=$2
currency=$3

if [ -z "$id" ] || [ -z "$price" ] || [ -z "$currency" ]; then
    echo "Usage: $0 <id> <price> <currency>"
    exit 1
fi

curl -X PUT -H "Content-Type: application/json" -d '{"value": '$price', "currency": "'$currency'"}' $SERVER_ADDR/products/$id

