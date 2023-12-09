package main

import (
	"congestion-calculator/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

const h = "localhost:8080"

func main() {
	r := SetupServer()
	fmt.Println("Listening on", h)
	r.Run(h)
}

func SetupServer() *gin.Engine {
	r := gin.Default()
	r.Use(gin.ErrorLogger())
	r.GET("/", ServeDocs)
	r.POST("/congestion-calculator/:location", controllers.HandleCongestionCalculation)

	return r
}

func ServeDocs(c *gin.Context) {
	c.File("./APIREADME.md")
}
