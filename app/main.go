package main

import (
	"github.com/0l1v3rr/gourly/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/api/v1/urls", controller.GetAllUrls)
	r.GET("/api/v1/urls/:id", controller.GetUrlById)
	r.POST("/api/v1/urls", controller.CreateUrl)
	r.PATCH("/api/v1/urls/:id", controller.IncrementUrlClicks)
	r.DELETE("/api/v1/urls/:id", controller.DeleteUrl)

	r.Run()
}
