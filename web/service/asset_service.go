package service

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/zyp461476492/docker-app/database"
	"github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/utils"
)

func AddAsset(asset types.DockerAsset) types.RetMsg {
	db, err := database.GetBoltDB(utils.Config)
	if err != nil {
		return types.RetMsg{Res: false, Info: types.DATABASE_FAIL}
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("DockerAsset"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}
