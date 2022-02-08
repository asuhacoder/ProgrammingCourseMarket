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

type AuthRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func login(c *gin.Context) {
	conn, err := grpc.Dial(authAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbAuth.NewAuthClient(conn)

	var s AuthRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Login(ctx, &pbAuth.LoginRequest{
		Email:    s.Email,
		Password: s.Password,
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

func authRouters(router *gin.RouterGroup) {
	a := router.Group("/auth")
	a.POST("", login)
}
