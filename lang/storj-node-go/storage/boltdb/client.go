package boltdb

import (
	"time"

	"github.com/boltdb/bolt"
)

var defaltTimeout = 1 * time.Second

// Client is the implemented storage interface for the Bolt database
type Client struct {
	UsersBucket *bolt.Bucket
}

// NewBoltDB instantiates an instance of a new Bolt DB client
func NewBoltDB() (*Client, error) {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: defaltTimeout})
	if err != nil {
		return nil, err
	}

	b := &bolt.Bucket{}
	err = db.Update(func(tx *bolt.Tx) error {
		b, err = tx.CreateBucket([]byte("users"))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &Client{
		UsersBucket: b,
	}, nil
}
