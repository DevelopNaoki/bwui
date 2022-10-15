package server

import (
	//"net/http"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	engine.Static("/css", "src/css")
	engine.Static("/js", "src/js")
}
