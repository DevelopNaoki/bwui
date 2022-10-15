package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	engine.Static("/css", "view/css")
	engine.Static("/js", "view/js")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
