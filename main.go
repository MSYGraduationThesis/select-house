package main

import (
	"github.com/gin-gonic/gin"
	"select-house/handel"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/param", handel.GetHouseModelService)
	router.POST("/sum", handel.PostHouseModelService)

	router.Run(":80")
}
