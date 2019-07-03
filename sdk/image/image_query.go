package image

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

/**
返回 client 客户端查询的所有 image 信息
*/
func List(cli *client.Client) ([]types.ImageSummary, error) {
	ctx := context.Background()
	return cli.ImageList(ctx, types.ImageListOptions{})
}

/**
返回 id 对应 image 的历史信息
*/
func History(cli *client.Client, id string) ([]types.ImageHistory, error) {
	ctx := context.Background()
	return cli.ImageHistory(ctx, id)
}
