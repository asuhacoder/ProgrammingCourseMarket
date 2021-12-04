package gateway

import (
	"context"
	"log"
	"time"

	pbUser "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	userAddress = "user:50051"
)

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
