package gateway

import (
	"context"
	"io"
	"log"
	"time"

	pbCourse "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/course"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	courseAddress = "course:50053"
)

type CourseListRequest struct {
	UserId     string `form:"user_id" json:"user_id"`
	OnlyPublic bool   `form:"only_public" json:"only_public"`
	OnlyMine   bool   `form:"only_mine" json:"only_mine"`
}

type CourseCreateUpdateRequest struct {
	UserId       string `form:"user_id" json:"user_id"`
	Title        string `from:"title" json:"title"`
	Introduction string `from:"introduction" json:"introduction"`
	Image        string `from:"image" json:"image"`
	Price        int64  `from:"price" json:"price"`
	IsPublic     bool   `form:"is_public" json:"is_public"`
}

type CourseDeleteRequest struct {
	UserId string `form:"user_id" json:"user_id"`
}

func listCourses(c *gin.Context) {
	conn, err := grpc.Dial(courseAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCourse.NewCourseClient(conn)

	var s CourseListRequest
	err = c.ShouldBind(&s)
	if err != nil {
		log.Printf("failed to bind request: %v", err)
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.ListCourses(ctx, &pbCourse.ListCoursesRequest{
		UserId:     s.UserId,
		OnlyPublic: s.OnlyPublic,
		OnlyMine:   s.OnlyMine,
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
	conn, err := grpc.Dial(courseAddress, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("connected grpc server")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCourse.NewCourseClient(conn)

	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetCourse(ctx, &pbCourse.GetCourseRequest{
		Uuid: uuid,
	})

	if err != nil {
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
	conn, err := grpc.Dial(courseAddress, grpc.WithInsecure(), grpc.WithBlock())
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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
	conn, err := grpc.Dial(courseAddress, grpc.WithInsecure(), grpc.WithBlock())
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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
	conn, err := grpc.Dial(courseAddress, grpc.WithInsecure(), grpc.WithBlock())
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.DeleteCourse(ctx, &pbCourse.DeleteCourseRequest{
		UserId: s.UserId,
		Uuid:   uuid,
	})

	if err != nil {
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
