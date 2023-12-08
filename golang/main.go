package main

import (
	"congestion-calculator/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

const h = "localhost:8080"

func main() {
	fmt.Println("Listening on", h)
	r := gin.Default()
	r.Use(gin.ErrorLogger())
	r.GET("/", ServeDocs)
	r.POST("/congestion-calculator/:location", controllers.HandleCongestionCalculation)

	r.Run(h)
}

func ServeDocs(c *gin.Context) {
	c.File("./APIREADME.md")
}
