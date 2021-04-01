package controller

import (
	shwscontroller "shellme/controller/shws"

	"github.com/gin-gonic/gin"
)

// SetupRoutes ...
func SetupRoutes(r *gin.Engine) {

	// shell websocket
	shws := r.Group("/shws")
	shws.GET("/", shwscontroller.WsHandler)

}
