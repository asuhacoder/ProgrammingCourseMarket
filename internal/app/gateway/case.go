package gateway

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	pbCase "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/case"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var caseHost = os.Getenv("USER_HOST")
var caseAddress = caseHost + ":50056"

type CreateCaseRequest struct {
	Token    string `form:"token" json:"token"`
	LessonID string `form:"lesson_id" json:"lesson_id"`
	Input    string `form:"input" json:"input"`
	Output   string `form:"output" json:"output"`
}

type UpdateCaseRequest struct {
	Token    string `form:"token" json:"token"`
	Uuid     string `form:"uuid" json:"uuid"`
	LessonID string `form:"lesson_id" json:"lesson_id"`
	Input    string `form:"input" json:"input"`
	Output   string `form:"output" json:"output"`
}

type DeleteCaseRequest struct {
	Token string `form:"token" json:"token"`
}

func listCases(c *gin.Context) {
	conn, err := grpc.Dial(caseAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCase.NewCaseClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	lessonID := c.Query("lesson_id")
	stream, err := client.ListCases(ctx, &pbCase.ListCasesRequest{LessonId: lessonID})
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
			log.Fatalf("failed to get case stream: %v", nil)
			c.AbortWithStatus(400)
		} else {
			testCase := gin.H{
				"uuid":      r.GetUuid(),
				"user_id":   r.GetUserId(),
				"lesson_id": r.GetLessonId(),
				"input":     r.GetInput(),
				"output":    r.GetOutput(),
			}
			responces = append(responces, testCase)
		}
	}
	c.JSON(200, gin.H{
		"cases": responces,
	})
}

func getCase(c *gin.Context) {
	conn, err := grpc.Dial(caseAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCase.NewCaseClient(conn)

	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.GetCase(ctx, &pbCase.GetCaseRequest{
		Uuid: uuid,
	})

	if err != nil {
		log.Printf("failed to get case: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":      r.GetUuid(),
			"user_id":   r.GetUserId(),
			"lesson_id": r.GetLessonId(),
			"input":     r.GetInput(),
			"output":    r.GetOutput(),
		})
	}
}

func createCase(c *gin.Context) {
	log.Println("createCase func started")
	conn, err := grpc.Dial(caseAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCase.NewCaseClient(conn)

	var s CreateCaseRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.CreateCase(ctx, &pbCase.CreateCaseRequest{
		Token:    s.Token,
		LessonId: s.LessonID,
		Input:    s.Input,
		Output:   s.Output,
	})

	if err != nil {
		log.Printf("failed to create case: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":      r.GetUuid(),
			"user_id":   r.GetUserId(),
			"lesson_id": r.GetLessonId(),
			"input":     r.GetInput(),
			"output":    r.GetOutput(),
		})
	}
}

func updateCase(c *gin.Context) {
	log.Println("updateCase func started")
	conn, err := grpc.Dial(caseAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCase.NewCaseClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var s UpdateCaseRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	uuid := c.Param("uuid")

	r, err := client.UpdateCase(ctx, &pbCase.UpdateCaseRequest{
		Token:    s.Token,
		Uuid:     uuid,
		LessonId: s.LessonID,
		Input:    s.Input,
		Output:   s.Output,
	})

	if err != nil {
		log.Printf("failed to update case: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":      r.GetUuid(),
			"user_id":   r.GetUserId(),
			"lesson_id": r.GetLessonId(),
			"input":     r.GetInput(),
			"output":    r.GetOutput(),
		})
	}
}

func deleteCase(c *gin.Context) {
	log.Println("deleteCase func started")
	conn, err := grpc.Dial(caseAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCase.NewCaseClient(conn)

	var s DeleteCaseRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}
	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.DeleteCase(ctx, &pbCase.DeleteCaseRequest{
		Token: s.Token,
		Uuid:  uuid,
	})

	if err != nil {
		log.Printf("failed to delete case: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r)
	}
}

func caseRouters(router *gin.RouterGroup) {
	c := router.Group("/cases")
	c.GET("", listCases)
	c.GET(":uuid", getCase)
	c.POST("", createCase)
	c.PUT(":uuid", updateCase)
	c.DELETE(":uuid", deleteCase)
}
