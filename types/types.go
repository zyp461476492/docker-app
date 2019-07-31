package types

import "time"

type DockerAsset struct {
	Id         int    `storm:"id,increment" json:"id"`
	Ip         string `storm:"index" json:"ip"`
	Name       string `storm:"index" json:"name"`
	Port       int    `json:"port"`
	Version    string `json:"version"`
	CreateTime string
}

type Config struct {
	FileLocation string
	Timeout      time.Duration
}

type RetMsg struct {
	Res  bool
	Info string
	Obj  interface{}
}
