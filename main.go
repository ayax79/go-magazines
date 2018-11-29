package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ayax79/go-magazines/dao"
	"github.com/ayax79/go-magazines/service"

	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/server"
)

func main() {
	log.Printf("Starting go-magazines")
	var cfg *service.Config

	log.Print("Loading ./config.json")
	config.LoadJSONFile("./config.json", &cfg)

	log.Printf("Beginning server initialization")
	server.Init("magazines-service", cfg.Server)

	redisServer := os.Getenv("REDIS_SERVER")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisServerPort := fmt.Sprintf("%s:%s", redisServer, redisPort)
	redisConfig := dao.NewRedisConfig(redisServerPort, redisPassword, 0)
	log.Printf("Redis configuration %#v", redisConfig)

	log.Printf("Instantiating service instance")
	service, err := service.NewMagazineService(cfg, redisConfig)

	if err != nil {
		server.Log.Fatal("unable to create service: ", err)
	} else {
		log.Printf("Registering Service with server")
		err := server.Register(service)
		fmt.Printf("Starting magazine service on port: %v", cfg.Server.HTTPPort)
		if err != nil {
			server.Log.Fatal("unable to register service: ", err)
		}
	}

}
