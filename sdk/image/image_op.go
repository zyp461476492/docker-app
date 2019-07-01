package image

import (
	"context"
	"log"

	"github.com/docker/docker/client"
)

/**
修改 id 对应 image 的 tag
*/
func changeTag(cli *client.Client, id, tag string) bool {
	ctx := context.Background()
	err := cli.ImageTag(ctx, id, tag)

	if err != nil {
		log.Fatalf("id:%v update tag fail, %v", id, err)
		return false
	}

	return true
}
