package main

import (
	"log"
	"webtech/prototype/db"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db.Init()

	r := gin.Default()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	r.Use(c)
	registerRoutes(r)
	log.Fatal(r.Run(":8080"))
}
