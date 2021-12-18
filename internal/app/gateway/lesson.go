package gateway

import (
	"context"
	"io"
	"log"
	"reflect"
	"time"

	pbLesson "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/lesson"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	lessonAddress = "lesson:50054"
)

type LessonRequest struct {
	Token        string           `form:"token" json:"token"`
	CourseID     string           `form:"course_id" json:"course_id"`
	Title        string           `form:"title" json:"title"`
	Introduction string           `form:"introduction" json:"introduction"`
	Body         string           `form:"body" json:"body"`
	DefaultCode  string           `form:"default_code" json:"default_code"`
	AnswerCode   string           `form:"answer_code" json:"answer_code"`
	TestCase     []*pbLesson.Case `form:"test_case" json:"test_case"`
	Language     string           `form:"language" json:"language"`
}

func listLessons(c *gin.Context) {
	conn, err := grpc.Dial(lessonAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbLesson.NewLessonClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	courseID := c.Query("course_id")
	stream, err := client.ListLessons(ctx, &pbLesson.ListLessonsRequest{CourseId: courseID})
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
				"uuid":         r.GetUuid(),
				"user_id":      r.GetUserId(),
				"course_id":    r.GetCourseId(),
				"title":        r.GetTitle(),
				"introduction": r.GetIntroduction(),
				"body":         r.GetBody(),
				"default_code": r.GetDefaultCode(),
				"answer_code":  r.GetAnswerCode(),
				"test_case":    r.GetTestCase(),
				"language":     r.GetLanguage(),
			}
			responces = append(responces, lesson)
		}
	}
	c.JSON(200, gin.H{
		"lessons": responces,
	})
}

func getLesson(c *gin.Context) {
	conn, err := grpc.Dial(lessonAddress, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbLesson.NewLessonClient(conn)

	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetLesson(ctx, &pbLesson.GetLessonRequest{
		Uuid: uuid,
	})
	log.Println("got data")
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":         r.GetUuid(),
			"user_id":      r.GetUserId(),
			"course_id":    r.GetCourseId(),
			"title":        r.GetTitle(),
			"introduction": r.GetIntroduction(),
			"body":         r.GetBody(),
			"default_code": r.GetDefaultCode(),
			"answer_code":  r.GetAnswerCode(),
			"test_case":    r.GetTestCase(),
			"language":     r.GetLanguage(),
		})
	}
}

func createLesson(c *gin.Context) {
	log.Println("createLesson func started")
	conn, err := grpc.Dial(lessonAddress, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbLesson.NewLessonClient(conn)

	var s LessonRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind queries: %v", err)
		c.AbortWithStatus(400)
	}
	log.Printf("request: %v", s)
	log.Printf("s.TestCase: %v", s.TestCase)
	log.Printf("s.TestCase's type: %v", reflect.TypeOf(s.TestCase[0]))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.CreateLesson(ctx, &pbLesson.CreateLessonRequest{
		Token:        s.Token,
		CourseId:     s.CourseID,
		Title:        s.Title,
		Introduction: s.Introduction,
		Body:         s.Body,
		DefaultCode:  s.DefaultCode,
		AnswerCode:   s.AnswerCode,
		TestCase:     s.TestCase,
		Language:     s.Language,
	})
	log.Println("got data")
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":         r.GetUuid(),
			"user_id":      r.GetUserId(),
			"course_id":    r.GetCourseId(),
			"title":        r.GetTitle(),
			"introduction": r.GetIntroduction(),
			"body":         r.GetBody(),
			"default_code": r.GetDefaultCode(),
			"answer_code":  r.GetAnswerCode(),
			"test_case":    r.GetTestCase(),
			"language":     r.GetLanguage(),
		})
	}
}

func updateLesson(c *gin.Context) {
	log.Println("updateLesson func started")
	conn, err := grpc.Dial(lessonAddress, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbLesson.NewLessonClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var s LessonRequest
	c.Bind(&s)
	log.Printf("request: %v", s)

	r, err := client.UpdateLesson(ctx, &pbLesson.UpdateLessonRequest{
		Token:        s.Token,
		Title:        s.Title,
		Introduction: s.Introduction,
		Body:         s.Body,
		DefaultCode:  s.DefaultCode,
		AnswerCode:   s.AnswerCode,
		TestCase:     s.TestCase,
		Language:     s.Language,
	})
	log.Println("got data")
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":         r.GetUuid(),
			"user_id":      r.GetUserId(),
			"course_id":    r.GetCourseId(),
			"title":        r.GetTitle(),
			"introduction": r.GetIntroduction(),
			"body":         r.GetBody(),
			"default_code": r.GetDefaultCode(),
			"answer_code":  r.GetAnswerCode(),
			"test_case":    r.GetTestCase(),
			"language":     r.GetLanguage(),
		})
	}
}

func deleteLesson(c *gin.Context) {
	log.Println("deleteLesson func started")
	conn, err := grpc.Dial(lessonAddress, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbLesson.NewLessonClient(conn)

	token := c.Query("token")
	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.DeleteLesson(ctx, &pbLesson.DeleteLessonRequest{
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

func lessonRouters(router *gin.RouterGroup) {
	l := router.Group("/lessons")
	l.GET("", listLessons)
	l.GET(":uuid", getLesson)
	l.POST("", createLesson)
	l.PUT(":uuid", updateLesson)
	l.DELETE(":uuid", deleteLesson)
}
