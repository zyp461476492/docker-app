package router

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zyp461476492/docker-app/sdk/image"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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
	err := r.ParseForm()
	c, err := upgrader.Upgrade(w, r, nil)
	defer c.Close()
	if err != nil {
		log.Printf("err %s", err.Error())
	}
	assetId, err := strconv.Atoi(r.Form.Get("assetId"))
	if err != nil {
		log.Fatalln(err)
	}
	imageId := r.Form.Get("imageId")
	fmt.Printf("%d %s", assetId, imageId)
	res, stream := image.PullImage(2, "docker.io/library/redis")
	if !res.Res {
		c.WriteMessage(1, []byte("pull error"))
		return
	}

	reader := bufio.NewReader(stream)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read Error: %s", err)
			return
		}
		fmt.Println(str)
		//c.WriteMessage(1, []byte(str))
		c.WriteJSON(str)
	}

}
