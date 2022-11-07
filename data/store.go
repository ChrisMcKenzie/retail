package data

import (
	"encoding/json"
	"errors"

	bolt "go.etcd.io/bbolt"
)

var (
	ErrNotFound = errors.New("not found")
)

const (
	PRICE_BUCKET = "prices"
)

type Store struct {
	// connection to bolt db
	db *bolt.DB
}

func NewStore(path string) *Store {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		panic(err)
	}

	// create bucket for prices
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(PRICE_BUCKET))
		return err
	})

	if err != nil {
		panic(err)
	}

	return &Store{db: db}
}

func (s *Store) FindPrice(id string) (Price, error) {
	var price Price
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(PRICE_BUCKET))
		if b == nil {
			return ErrNotFound
		}
		r := b.Get([]byte(id))
		if r == nil {
			return ErrNotFound
		}

		if err := json.Unmarshal(r, &price); err != nil {
			return err
		}

		return nil
	})
	return price, err
}

func (s *Store) ChangePrice(id string, price Price) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("prices"))

		// encode price to json
		data, err := json.Marshal(price)
		if err != nil {
			return err
		}

		return b.Put([]byte(id), data)
	})
}

func (s *Store) Close() error {
	return s.db.Close()
}
