package image

import (
	"context"
	"github.com/docker/docker/client"
)

/**
修改 id 对应 image 的 tag
*/
func changeTag(cli *client.Client, id, tag string) error {
	ctx := context.Background()
	return cli.ImageTag(ctx, id, tag)
}
