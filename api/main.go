package main

import (
	"context"
	"log"
	"time"

	pbUser "github.com/Asuha-a/ProgrammingCourseMarket/api/pb/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	userAddress = "user:50051"
)

func login(c *gin.Context) {
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewAuthClient(conn)

	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Login(ctx, &pbUser.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r.GetToken())
	}
}

func signup(c *gin.Context) {
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewAuthClient(conn)

	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Signup(ctx, &pbUser.SignupRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r.GetToken())
	}
}

func userRouters(router *gin.RouterGroup) {
	u := router.Group("/users")
	u.POST("login", login)
	u.POST("signup", signup)
}

func runServer() {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	userRouters(v1)
	r.Run()
}

func main() {
	runServer()
}
