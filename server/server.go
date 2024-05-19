package server

import (
	"fmt"

	"github.com/Ben352/autosuggestApi/internal/config"
	"github.com/Ben352/autosuggestApi/internal/handler"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	port := config.AppConfig.Server.Port
	apiKey := config.AppConfig.Server.ApiKey
	handler.RegisterRoutes(r, apiKey)

	r.Run(fmt.Sprintf(":%d", port))
}
