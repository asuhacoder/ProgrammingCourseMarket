package gateway

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	userRouters(v1)
	authRouters(v1)
	courseRouters(v1)
	lessonRouters(v1)
	runnerRouters(v1)
	err := r.Run()
	if err != nil {
		log.Printf("failed to run gateway: %v", err)
	}
}
