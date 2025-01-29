package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Ping recieved to Pong Railway!"})
}