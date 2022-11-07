package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	client  *http.Client
	baseURL string
	key     string
}

func NewClient(baseURL string, key string) *Client {
	return &Client{
		client:  http.DefaultClient,
		key:     key,
		baseURL: baseURL,
	}
}

func (c *Client) FindProduct(id string) (*Product, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/redsky_aggregations/v1/redsky/case_study_v1?key=%s&tcin=%s", c.baseURL, c.key, id))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var product Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}
