package main

import (
	"Go-RestfulAPI/config"
	"Go-RestfulAPI/pkg/db"
	"Go-RestfulAPI/pkg/logging"
	"Go-RestfulAPI/router"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-26 14:57
* @Description:
**/

var (
	cfg  = pflag.StringP("config", "c", "", "config file path")
	port string
)

func main() {
	pflag.Parse()

	//init config
	if err := config.Init(*cfg); err != nil {
		logging.GetLogger().Error("init config error.", zap.Error(err))
		panic(err)
	}
	//init log
	logging.InitLog()
	//init db
	db.InitDB()

	port = viper.GetString("addr")

	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			logging.GetLogger().Error("The router has no response, or it might took too long to start up.", zap.Error(err))
		}
		logging.GetLogger().Info("The router has been deployed successfully.")
	}()

	logging.GetLogger().Info("Start to listening the incoming requests on http address" + port)
	log.Printf(http.ListenAndServe(port, g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		logging.GetLogger().Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
