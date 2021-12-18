package gateway

import (
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
	r.Run()
}
