package image

import (
	"context"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

/**
返回 client 客户端查询的所有 image 信息
*/
func listImage(cli *client.Client) []types.ImageSummary {
	ctx := context.Background()
	imageList, err := cli.ImageList(ctx, types.ImageListOptions{})

	if err != nil {
		log.Fatalf("listImage fail, %v", err)
		return nil
	}

	return imageList
}

/**
返回 id 对应 image 的历史信息
*/
func imageHistory(cli *client.Client, id string) []types.ImageHistory {
	ctx := context.Background()
	historyList, err := cli.ImageHistory(ctx, id)

	if err != nil {
		log.Fatalf("id:%v query history info fail, %v", id, err)
		return nil
	}

	return historyList
}
