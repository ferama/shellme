package main

import (
	"io/fs"
	"net/http"
	"shellme/controller"
	"shellme/ui"
	"shellme/utils"
	"shellme/wshub"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	flags := utils.GetFlags()

	hub := wshub.GetHubInstance()
	go hub.Run()

	if !*flags.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	// Creates a router without any middleware by default
	r := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Content-Type, Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	controller.SetupRoutes(r)

	// static files custom middleware
	// use the "build" dir (the webpack target) as static root
	fsRoot, _ := fs.Sub(ui.StaticFiles, "build")
	fileserver := http.FileServer(http.FS(fsRoot))
	r.Use(func(c *gin.Context) {
		fileserver.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	})

	log.Info("Starting up server at :8000")
	r.Run(":8000")
}
