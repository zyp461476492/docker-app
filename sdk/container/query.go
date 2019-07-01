package container

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"log"
)

func List(cli *client.Client) []types.Container {
	ctx := context.Background()
	list, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		log.Fatalf("list container fail, %v", err)
		return nil
	}

	return list
}

/**
container id or name
*/
func logs(cli *client.Client, container string) io.ReadCloser {
	ctx := context.Background()
	stream, err := cli.ContainerLogs(ctx, container, types.ContainerLogsOptions{})

	if err != nil {
		log.Fatalf("logs container fail, %v", err)
		return nil
	}

	return stream
}

func diff(cli *client.Client, id string) []types.ContainerChange {
	ctx := context.Background()
	changeList, err := cli.ContainerDiff(ctx, id)

	if err != nil {
		log.Fatalf("query diff container fail, %v", err)
		return nil
	}

	return changeList
}
