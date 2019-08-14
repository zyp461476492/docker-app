package database

import (
	"github.com/asdine/storm"
	"github.com/zyp461476492/docker-app/types"
	"go.etcd.io/bbolt"
	"log"
	"time"
)

func GetStorm(config types.Config) (*storm.DB, error) {
	return storm.Open(
		config.FileLocation, storm.BoltOptions(0600, &bbolt.Options{Timeout: config.Timeout * time.Second}))
}

func CloseStorm(db *storm.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("关闭数据库失败 %s", err.Error())
	}
}
