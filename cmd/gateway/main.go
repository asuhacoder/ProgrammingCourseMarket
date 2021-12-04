package main

import (
	"context"
	"log"
	"time"

	pbAuth "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/auth"
	pbUser "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	userAddress = "user:50051"
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
		c.JSON(200, r.GetToken())
	}
}

func createUser(c *gin.Context) {
	log.Println("createUser func started")
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewUserClient(conn)

	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.CreateUser(ctx, &pbUser.CreateUserRequest{
		Email:    email,
		Password: password,
	})
	log.Println("got data")
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r.GetToken())
	}
}

func userRouters(router *gin.RouterGroup) {
	u := router.Group("/users")
	u.POST("", createUser)
}

func authRouters(router *gin.RouterGroup) {
	a := router.Group("/auth")
	a.GET("", login)
}

func runServer() {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	userRouters(v1)
	authRouters(v1)
	r.Run()
}

func main() {
	runServer()
}
