package gateway

import (
	"context"
	"log"
	"time"

	pbAuth "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	authAddress = "auth:50052"
)

func login(c *gin.Context) {
	conn, err := grpc.Dial(authAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbAuth.NewAuthClient(conn)

	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Login(ctx, &pbAuth.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"token":      r.GetToken(),
			"uuid":       r.GetUuid(),
			"email":      r.GetEmail(),
			"permission": r.GetPermission(),
		})
	}
}

func AuthRouters(router *gin.RouterGroup) {
	a := router.Group("/auth")
	a.GET("", login)
}
