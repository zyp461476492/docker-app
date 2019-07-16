package router

import (
	"encoding/json"
	"github.com/zyp461476492/docker-app/sdk/client"
	"github.com/zyp461476492/docker-app/sdk/image"
	"github.com/zyp461476492/docker-app/types"
	"log"
	"net/http"
	"strconv"
)

func list(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 根据 ID 获取 docker 资源信息
	asset := types.DockerAsset{Id: id}
	cli, err := client.GetClient(asset)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	imageList, err := image.List(cli)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonByte, err := json.Marshal(imageList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func history(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	assetId, err := strconv.Atoi(r.Form.Get("assetId"))
	if err != nil {
		log.Fatalln(err)
	}

	imageId := r.Form.Get("imageId")
	// 根据 ID 获取 docker 资源信息
	asset := types.DockerAsset{Id: assetId}
	cli, err := client.GetClient(asset)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	history, err := image.History(cli, imageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonByte, err := json.Marshal(history)

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
