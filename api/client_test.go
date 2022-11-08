package api

import (
	"fmt"
	"testing"

	"github.com/jfrog/go-mockhttp"
)

var mockData = `
{
  "data": {
    "product": {
      "tcin": "13860428",
      "item": {
        "product_description": {
          "title": "The Big Lebowski (Blu-ray)",
          "downstream_description": "Jeff \"The Dude\" Lebowski (Bridges) is the victim of mistaken identity. Thugs break into his apartment in the errant belief that they are accosting Jeff Lebowski, the eccentric millionaire philanthropist, not the laid-back, unemployed Jeff Lebowski. In the aftermath, \"The Dude\" seeks restitution from his wealthy namesake. He and his buddies (Goodman and Buscemi) are swept up in a kidnapping plot that quickly spins out of control."
        },
        "enrichment": {
          "images": {
            "primary_image_url": "https://target.scene7.com/is/image/Target/GUEST_bac49778-a5c7-4914-8fbe-96e9cd549450"
          }
        },
        "product_classification": {
          "product_type_name": "ELECTRONICS",
          "merchandise_type_name": "Movies"
        },
        "primary_brand": {
          "name": "Universal Home Video"
        }
      }
    }
  }
}
`

func TestClientSuccess(t *testing.T) {
	// setup mockhttp client
	h := mockhttp.NewClient(mockhttp.NewClientEndpoint().When(
		mockhttp.Request().Method("GET").Path("/redsky_aggregations/v1/redsky/case_study_v1"),
	).Respond(
		mockhttp.Response().StatusCode(200).BodyString(mockData),
	))

	// create a new client
	c := &Client{
		client:  h.HttpClient(),
		baseURL: "http://localhost:8080",
		key:     "key",
	}

	// call the client
	p, err := c.FindProduct("1234")

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	fmt.Println(p)

	// validate the data is correct
	if p.Data.Product.TCIN != "13860428" {
		t.Errorf("unexpected tcin: %s", p.Data.Product.TCIN)
	}

	if p.Data.Product.Item.ProductDescription.Title != "The Big Lebowski (Blu-ray)" {
		t.Errorf("unexpected title: %s", p.Data.Product.Item.ProductDescription.Title)
	}
}

func TestClientError(t *testing.T) {
	// setup mockhttp client
	h := mockhttp.NewClient(mockhttp.NewClientEndpoint().When(
		mockhttp.Request().Method("GET").Path("/redsky_aggregations/v1/redsky/case_study_v1"),
	).Respond(
		mockhttp.Response().StatusCode(500),
	))

	// create a new client
	c := &Client{
		client:  h.HttpClient(),
		baseURL: "http://localhost:8080",
		key:     "key",
	}

	// call the client
	_, err := c.FindProduct("1234")

	if err == nil {
		t.Errorf("expected error")
	}
}
