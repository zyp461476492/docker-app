package types

import (
	"time"
)

type DockerAsset struct {
	Id         int    `storm:"id,increment" json:"id"`
	Ip         string `json:"ip"`
	AssetName  string `storm:"index" json:"assetName"`
	Port       int    `json:"port"`
	Version    string `json:"version"`
	CreateTime string `storm:"index"`
	Status     string `json:"status"`
}

type ContainerCreateInfo struct {
	AssetId       int       `json:"assetId"`
	ContainerName string    `json:"containerName"`
	ImageName     string    `json:"imageName"`
	PortList      []PortMap `json:"portList"`
}

type PortMap struct {
	Type       string `json:"type"`
	DockerPort int    `json:"dockerPort"`
	HostPort   int    `json:"hostPort"`
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
