package course

import (
	"context"
	"errors"
	"log"
	"net"

	db "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/db/course"
	pb "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/course"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	port = ":50053"
)

type server struct {
	pb.UnimplementedCourseServer
}

func (s *server) ListCourses(in *pb.ListCoursesRequest, stream pb.Course_ListCoursesServer) error {
	uuid := in.GetUserId()
	log.Println("ListCourses running")
	var courses []db.Course
	var result *gorm.DB
	if in.GetOnlyPublic() && in.GetOnlyMine() {
		result = db.DB.Where("IS_PUBLIC = ?", true).Where("USER_ID = ?", uuid).Find(&courses)
	} else if in.GetOnlyPublic() {
		result = db.DB.Where("IS_PUBLIC = ?", true).Find(&courses)
	} else if in.GetOnlyMine() {
		result = db.DB.Where("USER_ID = ?", uuid).Find(&courses)
	} else {
		result = db.DB.Find(&courses)
	}
	log.Println("listed courses")
	log.Println(result.Error)
	if result.Error != nil {
		log.Fatalf("failed to list courses: %v", result.Error)
		return result.Error
	}
	for _, course := range courses {
		if err := stream.Send(&pb.ListCoursesReply{
			Uuid:         course.UUID.String(),
			UserId:       course.USER_ID.String(),
			Title:        course.TITLE,
			Introduction: course.INTRODUCTION,
			Image:        course.IMAGE,
			Price:        int64(course.PRICE),
			IsPublic:     course.IS_PUBLIC,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetCourse(ctx context.Context, in *pb.GetCourseRequest) (*pb.GetCourseReply, error) {
	var course db.Course
	result := db.DB.First(&course, "UUID = ?", in.GetUuid())
	log.Println(course)
	if result.Error != nil {
		log.Printf("failed to get a course: %v", result.Error)
		return &pb.GetCourseReply{}, result.Error
	}
	return &pb.GetCourseReply{
		Uuid:         course.UUID.String(),
		UserId:       course.USER_ID.String(),
		Title:        course.TITLE,
		Introduction: course.INTRODUCTION,
		Image:        course.IMAGE,
		Price:        int64(course.PRICE),
		IsPublic:     course.IS_PUBLIC,
	}, nil
}

func (s *server) CreateCourse(ctx context.Context, in *pb.CreateCourseRequest) (*pb.CreateCourseReply, error) {
	uUID, err := uuid.FromString(in.GetUserId())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}

	course := db.Course{
		UUID:         uuid.Must(uuid.NewV4()),
		USER_ID:      uUID,
		TITLE:        in.GetTitle(),
		INTRODUCTION: in.GetIntroduction(),
		IMAGE:        in.GetImage(),
		PRICE:        int(in.GetPrice()),
		IS_PUBLIC:    in.GetIsPublic(),
	}
	log.Println(course)
	result := db.DB.Create(&course)
	if result.Error != nil {
		log.Printf("failed to create course: %v", result.Error)
		return &pb.CreateCourseReply{}, result.Error
	}
	CreatedAt := timestamppb.New(course.CREATED_AT)

	return &pb.CreateCourseReply{
		Uuid:         course.UUID.String(),
		UserId:       course.USER_ID.String(),
		Title:        course.TITLE,
		Introduction: course.INTRODUCTION,
		Image:        course.IMAGE,
		Price:        int64(course.PRICE),
		IsPublic:     course.IS_PUBLIC,
		CreatedAt:    CreatedAt,
	}, nil
}

func (s *server) UpdateCourse(ctx context.Context, in *pb.UpdateCourseRequest) (*pb.UpdateCourseReply, error) {
	var course db.Course
	userUuid, err := uuid.FromString(in.GetUserId())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}
	uUID, err := uuid.FromString(in.GetUuid())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}
	result := db.DB.First(&course, "UUID = ?", uUID)
	if result.Error != nil {
		log.Printf("failed to update course: %v", result.Error)
		return &pb.UpdateCourseReply{}, result.Error
	}
	if userUuid != course.USER_ID {
		err := errors.New("invalid user id")
		log.Printf("failed to update course: %v", err)
		return &pb.UpdateCourseReply{}, err
	}
	course.TITLE = in.GetTitle()
	course.INTRODUCTION = in.GetIntroduction()
	course.IMAGE = in.GetImage()
	course.PRICE = int(in.GetPrice())
	course.IS_PUBLIC = in.GetIsPublic()
	db.DB.Save(&course)

	return &pb.UpdateCourseReply{
		Uuid:         course.UUID.String(),
		UserId:       course.USER_ID.String(),
		Title:        course.TITLE,
		Introduction: course.INTRODUCTION,
		Image:        course.IMAGE,
		Price:        int64(course.PRICE),
		IsPublic:     course.IS_PUBLIC,
	}, nil
}

func (s *server) DeleteCourse(ctx context.Context, in *pb.DeleteCourseRequest) (*empty.Empty, error) {
	var course db.Course
	userUuid, err := uuid.FromString(in.GetUserId())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}
	uUID, err := uuid.FromString(in.GetUuid())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}
	result := db.DB.First(&course, "UUID = ?", uUID)
	if result.Error != nil {

		return new(empty.Empty), result.Error
	}
	if userUuid != course.USER_ID {
		return new(empty.Empty), errors.New("invalid user id")
	}
	result = db.DB.Delete(&course, "UUID = ?", uUID)
	log.Println(course, result.Error)
	if result.Error != nil {
		log.Printf("failed to delete a course: %v", result.Error)
		return new(empty.Empty), result.Error
	}
	return new(empty.Empty), nil
}

func RunServer() {

	db.Init()
	defer db.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCourseServer(s, &server{})
	log.Println("course grpc server running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
