package main

import (
	// 1 load first
	"seed/config"
	// 2
	"github.com/gin-gonic/gin"
	"net/http"
	"seed/api"
	"seed/middleware"
)

func main() {
	r := gin.New()
	r.StaticFS("/app", http.Dir("./app")).Use(middleware.AddStaticHeader())
	r.Use(middleware.AddHeader(), middleware.RecoveryWithWriter())
	r.StaticFS("/static", http.Dir("./upload"))
	api.NewApiServer(r.Group("api"))
	// Listen and serve on 0.0.0.0:8080
	r.Run(config.ServerConfig.String())
}
