package server

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/DevelopNaoki/bwui/config"
)

func StartServer(serverConfig config.ServerConfig) {
	engine := gin.Default()
	engine.LoadHTMLGlob("view/html/*.html")

	Router(engine)

	if serverConfig.SSL {
		engine.RunTLS(":"+strconv.Itoa(serverConfig.Port), serverConfig.ServerCertificate, serverConfig.PrivateKey)
	} else {
		engine.Run(":"+strconv.Itoa(serverConfig.Port))
	}
}
