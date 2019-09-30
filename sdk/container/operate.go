package container

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	myClient "github.com/zyp461476492/docker-app/sdk/client"
	myType "github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/web/service"
	"log"
	"time"
)

func Create(info myType.ContainerCreateInfo) myType.RetMsg {
	asset, err := service.GetAsset(info.AssetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	config := container.Config{
		Image: info.ImageName,
	}

	portMap := make(map[int]int)
	for _, portInfo := range info.PortList {
		count, ok := portMap[portInfo.DockerPort]
		if !ok {
			portMap[portInfo.DockerPort] = 0
			count = 0
		}
		portMap[portInfo.DockerPort] = count + 1
	}

	portBinding := make(map[nat.Port][]nat.PortBinding)
	for _, portInfo := range info.PortList {
		key := fmt.Sprintf("%d", portInfo.DockerPort) + "/" + portInfo.Type
		bindList, ok := portBinding[nat.Port(key)]
		if !ok {
			bindList = make([]nat.PortBinding, 0, portMap[portInfo.DockerPort])
		}
		bindInfo := nat.PortBinding{
			HostIP:   "0.0.0.0",
			HostPort: fmt.Sprintf("%d", portInfo.HostPort),
		}
		bindList = append(bindList, bindInfo)
		portBinding[nat.Port(key)] = bindList
	}

	hostConfig := container.HostConfig{
		PortBindings: portBinding,
	}

	networkConfig := network.NetworkingConfig{}
	body, err := cli.ContainerCreate(context.Background(), &config, &hostConfig, &networkConfig, info.ContainerName)
	if err != nil {
		log.Printf("创建失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Obj: body}
}

func Remove(assetId int, containerId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	err = cli.ContainerRemove(context.Background(), containerId, types.ContainerRemoveOptions{})
	if err != nil {
		log.Printf("容器删除失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true}
}

func Start(assetId int, containerId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	err = cli.ContainerStart(context.Background(), containerId, types.ContainerStartOptions{})
	if err != nil {
		log.Printf("启动失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Info: "启动成功"}
}

func Pause(assetId int, containerId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	err = cli.ContainerPause(context.Background(), containerId)
	if err != nil {
		log.Printf("暂停失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true}
}

func Unpause(assetId int, containerId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	err = cli.ContainerUnpause(context.Background(), containerId)
	if err != nil {
		log.Printf("恢复失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true}
}

func Stop(assetId int, containerId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	timeout := 2 * time.Second
	err = cli.ContainerStop(context.Background(), containerId, &timeout)
	if err != nil {
		log.Printf("停止失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true}
}
