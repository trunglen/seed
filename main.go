package main

import (
	"log"
	"os"
	// 1 load first
	_ "seed/config"
	// 2
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"seed/api"
	"seed/middleware"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
func main() {
	r := gin.New()
	r.StaticFS("/app", http.Dir("./app")).Use(middleware.AddStaticHeader())
	r.Use(middleware.AddHeader(), middleware.RecoveryWithWriter())
	r.StaticFS("/static", http.Dir("./upload"))
	api.NewApiServer(r.Group("api"))
	// Listen and serve on 0.0.0.0:8080
	var port, err = determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	r.Run(":" + port)
}
