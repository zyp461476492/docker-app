package router

import (
	"encoding/json"
	"github.com/zyp461476492/sync-app/sdk/client"
	"github.com/zyp461476492/sync-app/sdk/image"
	"github.com/zyp461476492/sync-app/types"
	"log"
	"net/http"
	"strconv"
)

func list(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		log.Fatalln(err)
	}
	// 根据 ID 获取 docker 资源信息
	asset := types.DockerAsset{Id: id}
	cli := client.GetClient(asset)

	jsonByte, err := json.Marshal(image.List(cli))

	if err != nil {
		log.Fatalln(err)
	}

	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func history(w http.ResponseWriter, r *http.Request) {
	assetId, err := strconv.Atoi(r.Form.Get("assetId"))
	if err != nil {
		log.Fatalln(err)
	}

	imageId := r.Form.Get("imageId")
	// 根据 ID 获取 docker 资源信息
	asset := types.DockerAsset{Id: assetId}
	cli := client.GetClient(asset)

	jsonByte, err := json.Marshal(image.History(cli, imageId))

	if err != nil {
		log.Fatalln(err)
	}

	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func ConfigRouter() {
	http.HandleFunc("/image/list", list)
	http.HandleFunc("/image/history", history)
}
