package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	globals "auth/globals"
	middleware "auth/middleware"
	routes "auth/routers"
	"auth/routers"
)


var router *gin.Engine

func main() {
	
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run("localhost:8080")

}
