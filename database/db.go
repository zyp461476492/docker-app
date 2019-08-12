package database

import (
	"github.com/asdine/storm"
	"github.com/zyp461476492/docker-app/types"
	"log"
)

var cli *storm.DB

func GetStorm(config *types.Config) (*storm.DB, error) {
	if cli == nil {
		cli, err := storm.Open(config.FileLocation)
		return cli, err
	}
	return cli, nil
}

func CloseStorm(db *storm.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("关闭数据库失败 %s", err.Error())
	}
}
