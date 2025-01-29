package routes

import (
	"github.com/gin-gonic/gin"
	"go_projects/api"
)


func SetupRoutes(r *gin.Engine){
	r.GET("/ping", api.PingHandler)
}