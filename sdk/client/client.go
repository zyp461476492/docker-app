package client

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	localType "github.com/zyp461476492/sync-app/types"
	"log"
	"strconv"
)

func GetClient(asset localType.DockerAsset) *client.Client {
	// tcp://192.168.184.123:2376
	host := "tcp://" + asset.Ip + ":" + strconv.Itoa(asset.Port)
	cli, err := client.NewClient(host, asset.Version, nil, nil)

	if err != nil {
		log.Fatalf("get client fail, %v", asset)
		return nil
	}

	return cli
}

func GetClientInfo(cli *client.Client) types.Info {
	ctx := context.Background()

	info, err := cli.Info(ctx)

	if err != nil {
		log.Fatalf("get client info fail")
		return types.Info{}
	}

	return info
}
