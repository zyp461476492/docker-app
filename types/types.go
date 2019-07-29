package types

import "time"

type DockerAsset struct {
	Id         int
	Ip         string
	Port       int
	Version    string
	CreateTime string
}

type Config struct {
	FileLocation string
	Timeout      time.Duration
}

type RetMsg struct {
	Res  bool
	Info string
}
