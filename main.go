package main
import (
	"github.com/gin-gonic/gin"
	"go_projects/routes"
)

func main(){
	router:= gin.Default()
	routes.SetupRoutes(router)
	router.Run("0.0.0.0:3030")
}