package router

import "net/http"

func ConfigRouter() {
	http.HandleFunc("/image/list", list)
	http.HandleFunc("/image/search", search)
	http.HandleFunc("/image/history", history)
	http.HandleFunc("/container/list", containerList)
	http.HandleFunc("/asset/add", addAsset)
	http.HandleFunc("/asset/update", updateAsset)
	http.HandleFunc("/asset/delete", deleteAsset)
	http.HandleFunc("/asset/list", listAsset)
	http.HandleFunc("/asset/info", dockerInfo)
}
