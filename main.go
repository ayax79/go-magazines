package main

import (
	"github.com/ayax79/go-magazines/service"

	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/server"
)

func main() {
	var cfg *service.Config
	config.LoadJSONFile("./config.json", &cfg)

	server.Init("magazines-json-proxy", cfg.Server)

	err := server.Register(service.NewMagazineService(cfg))
	if err != nil {
		server.Log.Fatal("unable to register service: ", err)
	}

}
