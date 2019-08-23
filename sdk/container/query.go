package container

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	myClient "github.com/zyp461476492/docker-app/sdk/client"
	myType "github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/web/service"
	"io"
	"log"
)

func List(id int) myType.RetMsg {
	asset, err := service.GetAsset(id)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	containerList, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Printf("查询失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Obj: containerList}
}

/**
container id or name
*/
func logs(cli *client.Client, container string) (io.ReadCloser, error) {
	ctx := context.Background()

	return cli.ContainerLogs(ctx, container, types.ContainerLogsOptions{})
}

func diff(cli *client.Client, id string) ([]types.ContainerChange, error) {
	ctx := context.Background()
	return cli.ContainerDiff(ctx, id)
}
