package router

import (
	"encoding/json"
	"github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/web/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func deleteAsset(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	info := r.Form.Get("info")
	idList := strings.Split(info, ",")

	assetList := []types.DockerAsset{}
	for _, s := range idList {
		id, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("err param %s", s)
		} else {
			asset := types.DockerAsset{Id: id}
			assetList = append(assetList, asset)
		}

	}

	msg := service.DeleteAsset(assetList)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func updateAsset(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	asset := types.DockerAsset{}
	err = json.Unmarshal([]byte(r.Form.Get("info")), &asset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := service.UpdateAsset(&asset)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func addAsset(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	asset := types.DockerAsset{}

	err = json.Unmarshal([]byte(r.FormValue("info")), &asset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := service.AddAsset(&asset)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}
