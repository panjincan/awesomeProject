package main

import (
	"awesomeProject/ginDemo/config"
	"awesomeProject/ginDemo/middleware"
	"awesomeProject/ginDemo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	engine.Use(middleware.LoggerToFile())
	router.InitRouter(engine)
	engine.Run(config.PORT)
}
