package main

import (
	"github.com/Ben352/autosuggestApi/internal/config"
	"github.com/Ben352/autosuggestApi/server"
)

func main() {
	config.LoadConfig()
	server.Run()
}
