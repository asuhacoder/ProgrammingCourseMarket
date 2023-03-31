package gateway

import (
	"context"
	"log"
	"os"
	"time"

	pbAuth "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var authHost = os.Getenv("USER_HOST")
var authAddress = authHost + ":50052"

type AuthnRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type AuthzRequest struct {
	Token string `form:"token" json:"token"`
}

func authn(c *gin.Context) {
	conn, err := grpc.Dial(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbAuth.NewAuthClient(conn)

	var s AuthnRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.Authn(ctx, &pbAuth.AuthnRequest{
		Email:    s.Email,
		Password: s.Password,
	})
	if err != nil {
		log.Printf("failed to authentication: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"token":        r.GetToken(),
			"uuid":         r.GetUuid(),
			"name":         r.GetName(),
			"introduction": r.GetIntroduction(),
			"email":        r.GetEmail(),
			"permission":   r.GetPermission(),
		})
	}
}

func authz(c *gin.Context) {
	conn, err := grpc.Dial(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbAuth.NewAuthClient(conn)

	var s AuthzRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.Authz(ctx, &pbAuth.AuthzRequest{
		Token: s.Token,
	})
	if err != nil {
		log.Printf("failed to authorization: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"token":        r.GetToken(),
			"uuid":         r.GetUuid(),
			"name":         r.GetName(),
			"introduction": r.GetIntroduction(),
			"email":        r.GetEmail(),
			"permission":   r.GetPermission(),
		})
	}
}

func authRouters(router *gin.RouterGroup) {
	an := router.Group("/authn")
	an.POST("", authn)
	az := router.Group("/authz")
	az.POST("", authz)
}
