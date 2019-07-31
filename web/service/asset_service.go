package service

import (
	"fmt"
	"github.com/asdine/storm"
	"github.com/zyp461476492/docker-app/database"
	"github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/utils"
	"time"
)

func AddAsset(asset *types.DockerAsset) types.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		return types.RetMsg{Res: false, Info: types.DATABASE_FAIL}
	}

	asset.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	asset.Status = "0"
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
		return types.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	var assetList []types.DockerAsset
	skip := 0
	if (page - 1) > 0 {
		skip = (page - 1) * pageSize
	}
	err = db.All(&assetList, storm.Limit(pageSize), storm.Skip(skip))
	if err != nil {
		return types.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	total, err := db.Count(&types.DockerAsset{})
	if err != nil {
		return types.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	obj := map[string]interface{}{}
	obj["list"] = assetList
	obj["total"] = total
	database.CloseStorm(db)
	return types.RetMsg{Res: true, Info: types.SUCCESS, Obj: obj}
}
