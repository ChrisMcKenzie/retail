package data

import (
	"os"
	"testing"
)

func TestStore(t *testing.T) {
	// setup new store
	store := NewStore("test.db")
	defer func() {
		store.Close()
		// remove test.db
		os.Remove("test.db")
	}()

	// add price data for product
	store.ChangePrice("1234", Price{Value: 12.0, CurrencyCode: USD})

	// get price data for product
	price, err := store.FindPrice("1234")

	if err != nil {
		t.Error(err)
	}

	if price.Value != 12.0 {
		t.Error("Price should be 12.0")
	}
}

func TestStoreError(t *testing.T) {
	// setup new store
	store := NewStore("test.db")
	defer func() {
		store.Close()
		// remove test.db
		os.Remove("test.db")
	}()

	// get price data for product
	_, err := store.FindPrice("1234")

	if err == nil {
		t.Error("Should return error")
	}
}
