package gateway

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "status OK",
		})
	})
	r.Use(cors.Default())
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
