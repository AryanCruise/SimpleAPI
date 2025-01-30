package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Ping recieved to Pong Railway!"})
}

func TestEnv(c *gin.Context){
	c.String(http.StatusOK, "Hello Environment variable : %s", os.Getenv("TEST_ENV"))
}

func HomePage(c *gin.Context){
	c.String(http.StatusOK, "Hello World!")
}