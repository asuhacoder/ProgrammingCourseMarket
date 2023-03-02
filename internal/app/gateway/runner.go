package gateway

import (
	"context"
	"log"
	"time"

	pbRunner "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/runner"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	runnerAddress = "runner:50055"
)

type RunnerRequest struct {
	Code     string `form:"code" json:"code"`
	Input    string `form:"input" json:"input"`
	Language string `form:"language" json:"language"`
	Version  string `form:"version" json:"version"`
}

func runCode(c *gin.Context) {
	log.Println("runCode func started")
	conn, err := grpc.Dial(runnerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbRunner.NewRunnerClient(conn)

	var s RunnerRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.RunCode(ctx, &pbRunner.RunCodeRequest{
		Code:     s.Code,
		Input:    s.Input,
		Language: s.Language,
		Version:  s.Version,
	})

	if err != nil {
		log.Printf("failed to run code: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"output": r.GetOutput(),
		})
	}
}

func runnerRouters(router *gin.RouterGroup) {
	l := router.Group("/runner")
	l.POST("", runCode)
}
