package gateway

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://skhole.club"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	api := r.Group("/api")
	v1 := api.Group("/v1")
	userRouters(v1)
	authRouters(v1)
	courseRouters(v1)
	lessonRouters(v1)
	runnerRouters(v1)
	caseRouters(v1)
	err := r.Run()
	if err != nil {
		log.Printf("failed to run gateway: %v", err)
	}
}
