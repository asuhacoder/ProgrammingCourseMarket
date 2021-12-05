package gateway

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	userRouters(v1)
	AuthRouters(v1)
	r.Run()
}