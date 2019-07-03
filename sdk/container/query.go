package container

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
)

func List(cli *client.Client) ([]types.Container, error) {
	ctx := context.Background()
	return cli.ContainerList(ctx, types.ContainerListOptions{})
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
