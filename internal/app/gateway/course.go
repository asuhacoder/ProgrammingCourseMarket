package gateway

import (
	"context"
	"io"
	"log"
	"strconv"
	"time"

	pbCourse "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/course"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	courseAddress = "course:50053"
)

func listCourses(c *gin.Context) {
	conn, err := grpc.Dial(courseAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbCourse.NewCourseClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.ListCourses(ctx, new(empty.Empty))
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
				"uuid":       r.GetUuid(),
				"user_id":    r.GetUserId(),
				"title":      r.GetTitle(),
				"introduction":  r.GetIntroduction(),
				"image":      r.GetImage(),
				"price":      r.GetPrice(),
				"published":  r.GetPublished(),
				"created_at": r.GetCreatedAt(),
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
	log.Println("got data")
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":       r.GetUuid(),
			"user_id":    r.GetUserId(),
			"title":      r.GetTitle(),
			"introduction":  r.GetIntroduction(),
			"image":      r.GetImage(),
			"price":      r.GetPrice(),
			"published":  r.GetPublished(),
			"created_at": r.GetCreatedAt(),
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

	token := c.Query("token")
	title := c.Query("title")
	introduction := c.Query("introduction")
	image := c.Query("image")
	price, err := strconv.ParseInt(c.Query("price"), 10, 64)
	if err != nil {
		c.AbortWithStatus(400)
	}
	published, err := strconv.ParseBool(c.Query("published"))
	if err != nil {
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.CreateCourse(ctx, &pbCourse.CreateCourseRequest{
		Token:     token,
		Title:     title,
		Introduction: introduction,
		Image:     image,
		Price:     price,
		Published: published,
	})
	log.Println("got data")
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":       r.GetUuid(),
			"user_id":    r.GetUserId(),
			"title":      r.GetTitle(),
			"introduction":  r.GetIntroduction(),
			"image":      r.GetImage(),
			"price":      r.GetPrice(),
			"published":  r.GetPublished(),
			"created_at": r.GetCreatedAt(),
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
	token := c.Query("token")
	title := c.Query("title")
	introduction := c.Query("introduction")
	image := c.Query("image")
	price, err := strconv.ParseInt(c.Query("price"), 10, 64)
	if err != nil {
		c.AbortWithStatus(400)
	}
	published, err := strconv.ParseBool(c.Query("published"))
	if err != nil {
		c.AbortWithStatus(400)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.UpdateCourse(ctx, &pbCourse.UpdateCourseRequest{
		Uuid:      uuid,
		Token:     token,
		Title:     title,
		Introduction: introduction,
		Image:     image,
		Price:     price,
		Published: published,
	})
	log.Println("got data")
	log.Println(err)
	if err != nil {
		c.AbortWithStatus(400)
	} else {
		c.JSON(200, gin.H{
			"uuid":       r.GetUuid(),
			"user_id":    r.GetUserId(),
			"title":      r.GetTitle(),
			"introduction":  r.GetIntroduction(),
			"image":      r.GetImage(),
			"price":      r.GetPrice(),
			"published":  r.GetPublished(),
			"created_at": r.GetCreatedAt(),
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

	token := c.Query("token")
	uuid := c.Param("uuid")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.DeleteCourse(ctx, &pbCourse.DeleteCourseRequest{
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

func courseRouters(router *gin.RouterGroup) {
	u := router.Group("/courses")
	u.GET("", listCourses)
	u.GET(":uuid", getCourse)
	u.POST("", createCourse)
	u.PUT(":uuid", updateCourse)
	u.DELETE(":uuid", deleteCourse)
}
