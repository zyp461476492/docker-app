package image

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/zyp461476492/docker-app/database"
	myClient "github.com/zyp461476492/docker-app/sdk/client"
	myType "github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/utils"
	"log"
)

/**
返回 client 客户端查询的所有 image 信息
*/
func List(id int) myType.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		log.Print(err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	asset := myType.DockerAsset{}
	err = db.One("Id", id, &asset)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	imageList, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.Printf("查询失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Obj: imageList}
}

/**
返回 id 对应 image 的历史信息
*/
func History(assetId int, imageId string) myType.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		log.Print(err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	asset := myType.DockerAsset{}
	err = db.One("Id", assetId, &asset)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	historyInfo, err := cli.ImageHistory(context.Background(), imageId)
	if err != nil {
		log.Printf("查询失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Obj: historyInfo}
}
