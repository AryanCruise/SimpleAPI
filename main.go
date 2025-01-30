package main

import (
	"go_projects/config"
	"go_projects/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	router:= gin.Default()
	routes.SetupRoutes(router)
	port := config.GetEnv("PORT", "8080")
	router.Run("0.0.0.0:" + port)
}