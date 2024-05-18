package server

import (
	"fmt"

	"github.com/Ben352/autosuggestApi/internal/config"
	"github.com/Ben352/autosuggestApi/internal/handler"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	handler.RegisterRoutes(r)
	port := config.AppConfig.Server.Port
	r.Run(fmt.Sprintf(":%d", port))
}
