package database

import (
	"github.com/boltdb/bolt"
	"github.com/zyp461476492/docker-app/types"
	"time"
)

func getBoltDB(config *types.Config) (*bolt.DB, error) {
	return bolt.Open(config.FileLocation,
		0600, &bolt.Options{Timeout: config.Timeout * time.Millisecond})
}

func CloseBoltDB(db *bolt.DB) error {
	return db.Close()
}
