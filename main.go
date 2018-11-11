package main

import (
	"fmt"
	"os"

	"github.com/ayax79/go-magazines/dao"
	"github.com/ayax79/go-magazines/service"

	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/server"
)

func main() {
	var cfg *service.Config
	config.LoadJSONFile("./config.json", &cfg)

	server.Init("magazines-json-proxy", cfg.Server)

	redisServer := os.Getenv("REDIS_SERVER")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisServerPort := fmt.Sprintf("%s:%s", redisServer, redisPort)
	redisConfig := dao.NewRedisConfig(redisServerPort, redisPassword, 0)

	service, err := service.NewMagazineService(cfg, redisConfig)

	if err != nil {
		err := server.Register(service)
		if err != nil {
			server.Log.Fatal("unable to register service: ", err)
		}
	}

}
