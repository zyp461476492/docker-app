package image

import (
	"github.com/zyp461476492/sync-app/sdk/client"
	"github.com/zyp461476492/sync-app/types"
	"testing"
)

func TestImageList(t *testing.T) {
	asset := types.DockerAsset{
		Ip:      "192.168.184.123",
		Port:    2376,
		Version: "v1.39",
	}

	cli := client.GetClient(asset)

	res := List(cli)

	t.Log(res)
}
