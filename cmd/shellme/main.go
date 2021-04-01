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

// // It will add all the files in ui/build, including hidden files.
// //go:embed ../../ui/build/*
// var staticFiles embed.FS

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
	fsRoot, _ := fs.Sub(ui.StaticFiles, "build")
	fileserver := http.FileServer(http.FS(fsRoot))
	r.Use(func(c *gin.Context) {
		fileserver.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	})

	log.Info("Starting up server at :8000")
	r.Run(":8000")
}
