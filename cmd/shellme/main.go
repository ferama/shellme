package main

import (
	"os"
	"shellme/controller"
	"shellme/utils"
	"shellme/wshub"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func main() {
	// log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	utils.GetFlags()

	hub := wshub.GetHubInstance()
	go hub.Run()

	// creates router with default middlewares
	// r := gin.Default()

	// Creates a router without any middleware by default
	r := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	//r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	isProduction := os.Getenv("PRODUCTION")
	// log.Printf("=== IsProduction: %v===\n", isProduction)
	// if isProduction != "1" {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Content-Type, Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// }

	controller.SetupRoutes(r)

	// static files
	if isProduction == "1" {
		r.Use(static.Serve("/", static.LocalFile("ui/build", true)))
	} else {
		r.GET("/", func(c *gin.Context) {
			c.Writer.WriteString("Dev Mode: run the ui using npm from the ui folder")
		})
	}

	// r.NoRoute(func(c *gin.Context) {
	// 	c.File("./ui/build")
	// })

	log.Info("Starting up server at :8000")
	r.Run(":8000")
}
