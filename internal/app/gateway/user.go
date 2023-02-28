package gateway

import (
	"context"
	"io"
	"log"
	"time"

	pbUser "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	userAddress = "user:50051"
)

type CreateUserRequest struct {
	Name         string `form:"name" json:"name"`
	Introduction string `form:"introduction" json:"introduction"`
	Email        string `form:"email" json:"email"`
	Password     string `form:"password" json:"password"`
}

type UpdateUserRequest struct {
	Token        string `form:"token" json:"token"`
	Uuid         string `form:"uuid" json:"uuid"`
	Name         string `form:"name" json:"name"`
	Introduction string `form:"introduction" json:"introduction"`
	Email        string `form:"email" json:"email"`
	Password     string `form:"password" json:"password"`
}

type DeleteUserRequest struct {
	Token string `form:"token" json:"token"`
}

func listUsers(c *gin.Context) {
	conn, err := grpc.Dial(userAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
				"uuid":         r.GetUuid(),
				"name":         r.GetName(),
				"introduction": r.GetIntroduction(),
				"email":        r.GetEmail(),
				"permission":   r.GetPermission(),
			}
			responces = append(responces, user)
		}
	}
	c.JSON(200, gin.H{
		"users": responces,
	})
}

func getUser(c *gin.Context) {
	conn, err := grpc.Dial(userAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	if err != nil {
		log.Printf("failed to get user: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":         r.GetUuid(),
			"name":         r.GetName(),
			"introduction": r.GetIntroduction(),
			"email":        r.GetEmail(),
			"permission":   r.GetPermission(),
		})
	}
}

func createUser(c *gin.Context) {
	log.Println("createUser func started")
	conn, err := grpc.Dial(userAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewUserClient(conn)

	var s CreateUserRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.CreateUser(ctx, &pbUser.CreateUserRequest{
		Name:         s.Name,
		Introduction: s.Introduction,
		Email:        s.Email,
		Password:     s.Password,
	})

	if err != nil {
		log.Printf("failed to create user: %v", err)
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

func updateUser(c *gin.Context) {
	log.Println("updateUser func started")
	conn, err := grpc.Dial(userAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewUserClient(conn)

	var s UpdateUserRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.UpdateUser(ctx, &pbUser.UpdateUserRequest{
		Token:        s.Token,
		Name:         s.Name,
		Introduction: s.Introduction,
		Email:        s.Email,
		Password:     s.Password,
		Uuid:         s.Uuid,
	})

	if err != nil {
		log.Printf("failed to update user: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r.GetToken())
	}
}

func deleteUser(c *gin.Context) {
	log.Println("deleteUser func started")
	conn, err := grpc.Dial(userAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUser.NewUserClient(conn)

	uuid := c.Param("uuid")
	var s DeleteUserRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.DeleteUser(ctx, &pbUser.DeleteUserRequest{
		Token: s.Token,
		Uuid:  uuid,
	})

	if err != nil {
		log.Printf("failed to delete user: %v", err)
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
