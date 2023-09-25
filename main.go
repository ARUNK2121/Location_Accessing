package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	engine.LoadHTMLGlob("templates/*.html")

	engine.Use(gin.Logger())

	engine.GET("/", indexHandler)
	engine.GET("/location", DisplayLocation)
}

func DisplayLocation(c *gin.Context) {

}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
