package gateway

import (
	"context"
	"io"
	"log"
	"time"

	pbUser "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	userAddress = "user:50051"
)

func listUsers(c *gin.Context) {
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewUserClient(conn)

	token := c.Query("token")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.ListUsers(ctx, &pbUser.ListUsersRequest{
		Token: token,
	})
	if err != nil {
		log.Printf("failed to access grpc server: %v", err)
	}
	var responces []gin.H
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to get user stream: %v", nil)
			c.AbortWithStatus(400)
		} else {
			user := gin.H{
				"uuid":       r.GetUuid(),
				"email":      r.GetEmail(),
				"permission": r.GetPermission(),
				"password":   r.GetPassword(),
			}
			responces = append(responces, user)
		}
	}
	c.JSON(200, gin.H{
		"users": responces,
	})
}

func getUser(c *gin.Context) {
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewUserClient(conn)

	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetUser(ctx, &pbUser.GetUserRequest{
		Uuid: uuid,
	})
	log.Println("got data")
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":       r.GetUuid(),
			"email":      r.GetEmail(),
			"permission": r.GetPermission(),
		})
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
		c.JSON(200, gin.H{
			"token":      r.GetToken(),
			"uuid":       r.GetUuid(),
			"email":      r.GetEmail(),
			"permission": r.GetPermission(),
		})
	}
}

func updateUser(c *gin.Context) {
	log.Println("updateUser func started")
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewUserClient(conn)

	token := c.Query(("token"))
	newEmail := c.Query("email")
	newPassword := c.Query("password")
	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.UpdateUser(ctx, &pbUser.UpdateUserRequest{
		Token:       token,
		NewEmail:    newEmail,
		NewPassword: newPassword,
		Uuid:        uuid,
	})
	log.Println("got data")
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r.GetToken())
	}
}

func deleteUser(c *gin.Context) {
	log.Println("deleteUser func started")
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewUserClient(conn)

	token := c.Query("token")
	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.DeleteUser(ctx, &pbUser.DeleteUserRequest{
		Token: token,
		Uuid:  uuid,
	})
	log.Println("got data")
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r)
	}
}

func userRouters(router *gin.RouterGroup) {
	u := router.Group("/users")
	u.GET("", listUsers)
	u.GET(":uuid", getUser)
	u.POST("", createUser)
	u.PUT(":uuid", updateUser)
	u.DELETE(":uuid", deleteUser)
}
