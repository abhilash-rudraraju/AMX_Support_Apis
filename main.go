package main

import (
	"amx_support_api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/mockPosition", handlers.MockPositionController)
	router.POST("/order/v2/placeOrder", handlers.ReverserOrderUat)
	router.Run("localhost:8020")
}

