package image

import (
	"github.com/zyp461476492/docker-app/sdk/client"
	"github.com/zyp461476492/docker-app/types"
	"testing"
)

func TestImageList(t *testing.T) {
	asset := types.DockerAsset{
		Ip:      "192.168.184.123",
		Port:    2376,
		Version: "v1.39",
	}

	cli, err := client.GetClient(asset)

	if err != nil {
		t.Fatal(err)
	}

	res, err := List(cli)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
