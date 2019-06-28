package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	//改成自己的ip 加端口
	cli, err := client.NewClient("tcp://192.168.184.123:2376", "v1.39", nil, nil)
	log(err)
	listImage(cli)
	//pullImage(cli)
}

// 列出镜像
func listImage(cli *client.Client) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	log(err)

	for index, image := range images {
		image.ID = "unit" + string(index) + ":" + image.ID
		v,err := json.Marshal(image)
		if err != nil {
			log(err)
		}
		fmt.Printf("%+v\n",string(v))
	}
}

func pullImage(cli *client.Client) {
	ctx := context.Background()
	out, err := cli.ImagePull(ctx, "docker.io/library/node:7.8.0-alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	value, err := io.Copy(os.Stdout, out)

	if err != nil {
		fmt.Println(value)
		log(err)
	}
}

func log(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}
}

