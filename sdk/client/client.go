package client

import (
	"github.com/docker/docker/client"
	"github.com/zyp461476492/sync-app/types"
	"log"
	"strconv"
)

func GetClient(asset types.DockerAsset) *client.Client {
	// tcp://192.168.184.123:2376
	host := "tcp://" + asset.Ip + ":" + strconv.Itoa(asset.Port)
	cli, err := client.NewClient(host, asset.Version, nil, nil)

	if err != nil {
		log.Fatalf("get client fail, %v", asset)
		return nil
	}

	return cli
}
