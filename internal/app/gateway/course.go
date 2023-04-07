package gateway

import (
	"context"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	pbCourse "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/course"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var courseHost = os.Getenv("COURSE_HOST")
var courseAddress = courseHost + ":50053"

type CourseCreateUpdateRequest struct {
	UserId       string `form:"user_id" json:"user_id"`
	Title        string `form:"title" json:"title"`
	Introduction string `form:"introduction" json:"introduction"`
	Image        string `form:"image" json:"image"`
	Price        int64  `form:"price" json:"price"`
	IsPublic     bool   `form:"is_public" json:"is_public"`
}

type CourseDeleteRequest struct {
	UserId string `form:"user_id" json:"user_id"`
}

func listCourses(c *gin.Context) {
	conn, err := grpc.Dial(courseAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCourse.NewCourseClient(conn)

	userId := c.Query("user_id")
	onlyPublicString := c.Query("only_public")
	onlyMineString := c.Query("only_mine")
	onlyPublic, err := strconv.ParseBool(onlyPublicString)
	if err != nil {
		log.Printf("failed to convert string to bool: %v", err)
	}
	onlyMine, err := strconv.ParseBool(onlyMineString)
	if err != nil {
		log.Printf("failed to convert string to bool: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stream, err := client.ListCourses(ctx, &pbCourse.ListCoursesRequest{
		UserId:     userId,
		OnlyPublic: onlyPublic,
		OnlyMine:   onlyMine,
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
			log.Fatalf("failed to get course stream: %v", nil)
			c.AbortWithStatus(400)
		} else {
			course := gin.H{
				"uuid":         r.GetUuid(),
				"user_id":      r.GetUserId(),
				"title":        r.GetTitle(),
				"introduction": r.GetIntroduction(),
				"image":        r.GetImage(),
				"price":        r.GetPrice(),
				"is_public":    r.GetIsPublic(),
				"created_at":   r.GetCreatedAt(),
			}
			responces = append(responces, course)
		}
	}
	c.JSON(200, gin.H{
		"courses": responces,
	})
}

func getCourse(c *gin.Context) {
	conn, err := grpc.Dial(courseAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCourse.NewCourseClient(conn)

	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.GetCourse(ctx, &pbCourse.GetCourseRequest{
		Uuid: uuid,
	})

	if err != nil {
		log.Printf("failed to get course: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":         r.GetUuid(),
			"user_id":      r.GetUserId(),
			"title":        r.GetTitle(),
			"introduction": r.GetIntroduction(),
			"image":        r.GetImage(),
			"price":        r.GetPrice(),
			"is_public":    r.GetIsPublic(),
			"created_at":   r.GetCreatedAt(),
		})
	}
}

func createCourse(c *gin.Context) {
	log.Println("createCourse func started")
	conn, err := grpc.Dial(courseAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCourse.NewCourseClient(conn)

	var s CourseCreateUpdateRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.CreateCourse(ctx, &pbCourse.CreateCourseRequest{
		UserId:       s.UserId,
		Title:        s.Title,
		Introduction: s.Introduction,
		Image:        s.Image,
		Price:        s.Price,
		IsPublic:     s.IsPublic,
	})

	if err != nil {
		log.Printf("failed to create course: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":         r.GetUuid(),
			"user_id":      r.GetUserId(),
			"title":        r.GetTitle(),
			"introduction": r.GetIntroduction(),
			"image":        r.GetImage(),
			"price":        r.GetPrice(),
			"is_public":    r.GetIsPublic(),
			"created_at":   r.GetCreatedAt(),
		})
	}
}

func updateCourse(c *gin.Context) {
	log.Println("updateCourse func started")
	conn, err := grpc.Dial(courseAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCourse.NewCourseClient(conn)

	uuid := c.Param("uuid")
	var s CourseCreateUpdateRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.UpdateCourse(ctx, &pbCourse.UpdateCourseRequest{
		Uuid:         uuid,
		UserId:       s.UserId,
		Title:        s.Title,
		Introduction: s.Introduction,
		Image:        s.Image,
		Price:        s.Price,
		IsPublic:     s.IsPublic,
	})

	if err != nil {
		log.Printf("failed to update course: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":         r.GetUuid(),
			"user_id":      r.GetUserId(),
			"title":        r.GetTitle(),
			"introduction": r.GetIntroduction(),
			"image":        r.GetImage(),
			"price":        r.GetPrice(),
			"is_public":    r.GetIsPublic(),
			"created_at":   r.GetCreatedAt(),
		})
	}
}

func deleteCourse(c *gin.Context) {
	log.Println("deleteCourse func started")
	conn, err := grpc.Dial(courseAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.WaitForReady(true)))
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCourse.NewCourseClient(conn)

	uuid := c.Param("uuid")
	var s CourseDeleteRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := client.DeleteCourse(ctx, &pbCourse.DeleteCourseRequest{
		UserId: s.UserId,
		Uuid:   uuid,
	})

	if err != nil {
		log.Printf("failed to delete course: %v", err)
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, r)
	}
}

func courseRouters(router *gin.RouterGroup) {
	c := router.Group("/courses")
	c.GET("", listCourses)
	c.GET(":uuid", getCourse)
	c.POST("", createCourse)
	c.PUT(":uuid", updateCourse)
	c.DELETE(":uuid", deleteCourse)
}
