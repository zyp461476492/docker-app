package router

import (
	"encoding/json"
	"github.com/zyp461476492/docker-app/sdk/image"
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
	id, err := strconv.Atoi(r.Form.Get("index"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := image.List(id)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	assetId, err := strconv.Atoi(r.Form.Get("assetId"))
	if err != nil {
		log.Fatalln(err)
	}
	term := r.Form.Get("term")

	msg := image.Search(assetId, term)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
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

	msg := image.History(assetId, imageId)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func imagePull(w http.ResponseWriter, r *http.Request) {

}
