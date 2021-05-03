package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.GET("/showSum", handleShowSum)

	router.Run(":9090")
}

func handleShowSum(ctx *gin.Context){
	ctx.JSON(http.StatusOK, Sum(8, 6))
}

func Sum(a, b int) int {
	return a + b
}