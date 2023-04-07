package gateway

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	pbLesson "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/lesson"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var lessonHost = os.Getenv("LESSON_HOST")
var lessonAddress = lessonHost + ":50054"

type CreateLessonRequest struct {
	Token          string `form:"token" json:"token"`
	CourseID       string `form:"course_id" json:"course_id"`
	SequenceNumber int64  `form:"sequence_number" json:"sequence_number"`
	Title          string `form:"title" json:"title"`
	Introduction   string `form:"introduction" json:"introduction"`
	Body           string `form:"body" json:"body"`
	DefaultCode    string `form:"default_code" json:"default_code"`
	AnswerCode     string `form:"answer_code" json:"answer_code"`
	Language       string `form:"language" json:"language"`
}

type UpdateLessonRequest struct {
	Token          string `form:"token" json:"token"`
	Uuid           string `form:"uuid" json:"uuid"`
	CourseID       string `form:"course_id" json:"course_id"`
	SequenceNumber int64  `form:"sequence_number" json:"sequence_number"`
	Title          string `form:"title" json:"title"`
	Introduction   string `form:"introduction" json:"introduction"`
	Body           string `form:"body" json:"body"`
	DefaultCode    string `form:"default_code" json:"default_code"`
	AnswerCode     string `form:"answer_code" json:"answer_code"`
	Language       string `form:"language" json:"language"`
}

type DeleteLessonRequest struct {
	Token string `form:"token" json:"token"`
}

func listLessons(c *gin.Context) {
	conn, err := grpc.Dial(lessonAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbLesson.NewLessonClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	courseID := c.Query("course_id")
	stream, err := client.ListLessons(ctx, &pbLesson.ListLessonsRequest{CourseId: courseID})
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
			log.Fatalf("failed to get lesson stream: %v", nil)
			c.AbortWithStatus(400)
		} else {
			lesson := gin.H{
				"uuid":            r.GetUuid(),
				"user_id":         r.GetUserId(),
				"course_id":       r.GetCourseId(),
				"sequence_number": r.GetSequenceNumber(),
				"title":           r.GetTitle(),
				"introduction":    r.GetIntroduction(),
				"body":            r.GetBody(),
				"default_code":    r.GetDefaultCode(),
				"answer_code":     r.GetAnswerCode(),
				"language":        r.GetLanguage(),
			}
			responces = append(responces, lesson)
		}
	}
	c.JSON(200, gin.H{
		"lessons": responces,
	})
}

func getLesson(c *gin.Context) {
	conn, err := grpc.Dial(lessonAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbLesson.NewLessonClient(conn)

	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.GetLesson(ctx, &pbLesson.GetLessonRequest{
		Uuid: uuid,
	})

	if err != nil {
		log.Printf("failed to get lesson: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":            r.GetUuid(),
			"user_id":         r.GetUserId(),
			"course_id":       r.GetCourseId(),
			"sequence_number": r.GetSequenceNumber(),
			"title":           r.GetTitle(),
			"introduction":    r.GetIntroduction(),
			"body":            r.GetBody(),
			"default_code":    r.GetDefaultCode(),
			"answer_code":     r.GetAnswerCode(),
			"language":        r.GetLanguage(),
		})
	}
}

func createLesson(c *gin.Context) {
	log.Println("createLesson func started")
	conn, err := grpc.Dial(lessonAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbLesson.NewLessonClient(conn)

	var s CreateLessonRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.CreateLesson(ctx, &pbLesson.CreateLessonRequest{
		Token:          s.Token,
		CourseId:       s.CourseID,
		SequenceNumber: s.SequenceNumber,
		Title:          s.Title,
		Introduction:   s.Introduction,
		Body:           s.Body,
		DefaultCode:    s.DefaultCode,
		AnswerCode:     s.AnswerCode,
		Language:       s.Language,
	})

	if err != nil {
		log.Printf("failed to create lesson: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":            r.GetUuid(),
			"user_id":         r.GetUserId(),
			"course_id":       r.GetCourseId(),
			"sequence_number": r.GetSequenceNumber(),
			"title":           r.GetTitle(),
			"introduction":    r.GetIntroduction(),
			"body":            r.GetBody(),
			"default_code":    r.GetDefaultCode(),
			"answer_code":     r.GetAnswerCode(),
			"language":        r.GetLanguage(),
		})
	}
}

func updateLesson(c *gin.Context) {
	log.Println("updateLesson func started")
	conn, err := grpc.Dial(lessonAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbLesson.NewLessonClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var s UpdateLessonRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	uuid := c.Param("uuid")

	r, err := client.UpdateLesson(ctx, &pbLesson.UpdateLessonRequest{
		Token:          s.Token,
		Uuid:           uuid,
		SequenceNumber: s.SequenceNumber,
		Title:          s.Title,
		Introduction:   s.Introduction,
		Body:           s.Body,
		DefaultCode:    s.DefaultCode,
		AnswerCode:     s.AnswerCode,
		Language:       s.Language,
	})

	if err != nil {
		log.Printf("failed to update lesson: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":            r.GetUuid(),
			"user_id":         r.GetUserId(),
			"course_id":       r.GetCourseId(),
			"sequence_number": r.GetSequenceNumber(),
			"title":           r.GetTitle(),
			"introduction":    r.GetIntroduction(),
			"body":            r.GetBody(),
			"default_code":    r.GetDefaultCode(),
			"answer_code":     r.GetAnswerCode(),
			"language":        r.GetLanguage(),
		})
	}
}

func deleteLesson(c *gin.Context) {
	log.Println("deleteLesson func started")
	conn, err := grpc.Dial(lessonAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbLesson.NewLessonClient(conn)

	var s DeleteLessonRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}
	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.DeleteLesson(ctx, &pbLesson.DeleteLessonRequest{
		Token: s.Token,
		Uuid:  uuid,
	})

	if err != nil {
		log.Printf("failed to delete lesson: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r)
	}
}

func lessonRouters(router *gin.RouterGroup) {
	l := router.Group("/lessons")
	l.GET("", listLessons)
	l.GET(":uuid", getLesson)
	l.POST("", createLesson)
	l.PUT(":uuid", updateLesson)
	l.DELETE(":uuid", deleteLesson)
}
