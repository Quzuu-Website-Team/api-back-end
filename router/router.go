package router

import (
	"github.com/gin-gonic/gin"
	"go-dp.abdanhafidz.com/config"
	"go-dp.abdanhafidz.com/controller"
)

func StartService() {
	router := gin.Default()
	routerGroup := router.Group("/api")
	{
		routerGroup.POST("/login", controller.LoginController)
		routerGroup.POST("/register", controller.RegisterController)
	}
	router.Run(config.TCP_ADDRESS)
}
