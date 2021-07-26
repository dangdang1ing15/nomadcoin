package db

import (
	"github.com/boltdb/bolt"
	"github.com/dangdang1ing15/nomadcoin/utils"
)

var db *bolt.DB

const (
	dbName      = "blockchaun.db"
	dataBucket  = "data"
	blockBucket = "blocks"
)

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		utils.HandleErr(err)
		db = dbPointer
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(data))
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blockBucket))
			return err
		})
		utils.HandleErr(err)
	}
	return db
}
