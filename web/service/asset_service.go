package service

import (
	"fmt"
	"github.com/asdine/storm"
	"github.com/zyp461476492/docker-app/database"
	"github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/utils"
)

func AddAsset(asset *types.DockerAsset) types.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		return types.RetMsg{Res: false, Info: types.DATABASE_FAIL}
	}

	err = db.Save(asset)

	if err != nil {
		return types.RetMsg{Res: false, Info: err.Error()}
	}

	database.CloseStorm(db)
	return types.RetMsg{Res: true}
}

func UpdateAsset(asset *types.DockerAsset) types.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		return types.RetMsg{Res: false, Info: types.DATABASE_FAIL}
	}

	err = db.Update(asset)

	if err != nil {
		return types.RetMsg{Res: false, Info: err.Error()}
	}

	database.CloseStorm(db)
	return types.RetMsg{Res: true}
}

func DeleteAsset(assetList []types.DockerAsset) types.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		return types.RetMsg{Res: false, Info: types.DATABASE_FAIL}
	}

	count := 0
	for _, asset := range assetList {
		err := db.DeleteStruct(&asset)
		if err != nil {
			count++
		}
	}

	info := fmt.Sprintf("成功数量 %d, 失败数量: %d", len(assetList)-count, count)

	database.CloseStorm(db)
	return types.RetMsg{Res: true, Info: info}
}

func ListAsset(page, pageSize int) types.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		return types.RetMsg{Res: false, Info: types.DATABASE_FAIL, Obj: nil}
	}
	var assetList []types.DockerAsset
	skip := 0
	if (page - 1) > 0 {
		skip = (page - 1) * pageSize
	}
	err = db.All(&assetList, storm.Limit(pageSize), storm.Skip(skip))
	if err != nil {
		return types.RetMsg{Res: false, Info: types.FAIL, Obj: nil}
	}

	database.CloseStorm(db)
	return types.RetMsg{Res: true, Info: types.SUCCESS, Obj: assetList}
}
