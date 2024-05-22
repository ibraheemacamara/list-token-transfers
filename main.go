package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ibraheemacamara/list-token-transfers/api"
	"github.com/ibraheemacamara/list-token-transfers/config"
	"github.com/ibraheemacamara/list-token-transfers/utils"
)

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading the config: ", err)
	}
}

func main() {
	cfg := config.GetConfig()

	// Init rpc provider client
	rppcProviderClient := utils.Init(cfg.RpcProvider(), time.Second*20)
	controller := api.NewController(rppcProviderClient)

	router := gin.Default()

	router.POST("/", controller.MainHandler())
	router.Run(cfg.Hostname() + ":" + cfg.Port())
}
