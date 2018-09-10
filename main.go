package main

import (
	"service"

	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/server"
)

func main() {
	var cfg *service.config
	config.LoadJSONFile("./config.json", &cfg)

	server.Init("magazines-json-proxy", cfg.Server)

	err := server.Register(service.NewJSONService(cfg))
	if err != nil {
		server.Log.Fatal("unable to register service: ", err)
	}

}
