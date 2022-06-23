package routers

import (
	"github.com/gin-gonic/gin"
	"auth/controller"

)


func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", controller.LoginGetHandler())
	g.POST("/login", controller.LoginPostHandler())
	g.GET("/", controller.IndexGetHandler())

}

func PrivateRoutes(g *gin.RouterGroup) {

	g.GET("/dashboard", controller.DashboardGetHandler())
	g.GET("/logout", controller.LogoutGetHandler())

}