package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ChrisMcKenzie/retail/api"
	"github.com/ChrisMcKenzie/retail/data"
	"github.com/gin-gonic/gin"
)

func main() {

	// get env vars
	clientAddr := os.Getenv("CLIENT_ADDR")
	if clientAddr == "" {
		fmt.Println("CLIENT_ADDR not set")
		os.Exit(1)
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		fmt.Println("CLIENT_SECRET not set")
		os.Exit(1)
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		fmt.Println("DB_PATH not set")
		os.Exit(1)
	}

	// setup client for redsky api
	client := api.NewClient(clientAddr, clientSecret)

	// setup local store
	store := data.NewStore(dbPath)
	defer store.Close()

	r := gin.Default()
	r.GET("/products/:id", func(c *gin.Context) {
		// get id from url
		id := c.Param("id")

		// find product by id
		product, err := client.FindProduct(id)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error: %s", err)
		}

		// get price from local data store
		price, err := store.FindPrice(id)
		if err != nil {
			if err == data.ErrNotFound {
				// 404 if we don't have a price
				c.String(http.StatusNotFound, "Error: %s", err)
				return
			}

			c.String(http.StatusInternalServerError, "Error: %s", err)
			return
		}

		// build response
		response := gin.H{
			"id":            product.Data.Product.TCIN,
			"name":          product.Data.Product.Item.ProductDescription.Title,
			"current_price": price,
		}

		c.JSON(http.StatusOK, response)
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		// get id from url
		id := c.Param("id")

		// get price from request body
		var price data.Price
		err := c.BindJSON(&price)

		// update price in local data store
		err = store.ChangePrice(id, price)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error: %s", err)
		}

		c.Status(http.StatusAccepted)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
