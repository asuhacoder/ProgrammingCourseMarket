package lesson

import (
	"context"
	"errors"
	"log"
	"net"

	db "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/db/lesson"
	jwt "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/jwt"
	pb "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/lesson"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port = ":50054"
)

type server struct {
	pb.UnimplementedLessonServer
}

func (s *server) ListLessons(rect *pb.ListLessonsRequest, stream pb.Lesson_ListLessonsServer) error {
	log.Println("ListLessons running")
	var lessons []db.Lesson
	result := db.LessonDB.Where("COURSE_ID = ?", rect.GetCourseId()).Find(&lessons)
	log.Println("got all lessons")
	log.Println(result.Error)
	if result.Error != nil {
		log.Fatalf("failed to list lessons: %v", result.Error)
		return result.Error
	}

	for _, lesson := range lessons {
		if err := stream.Send(&pb.ListLessonsReply{
			Uuid:           lesson.UUID.String(),
			UserId:         lesson.USER_ID.String(),
			CourseId:       lesson.COURSE_ID.String(),
			SequenceNumber: lesson.SEQUENCE_NUMBER,
			Title:          lesson.TITLE,
			Introduction:   lesson.INTRODUCTION,
			Body:           lesson.BODY,
			DefaultCode:    lesson.DEFAULT_CODE,
			AnswerCode:     lesson.ANSWER_CODE,
			Language:       lesson.LANGUAGE,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetLesson(ctx context.Context, in *pb.GetLessonRequest) (*pb.GetLessonReply, error) {
	var lesson db.Lesson
	result := db.LessonDB.First(&lesson, "UUID = ?", in.GetUuid())
	log.Println(lesson)
	if result.Error != nil {
		log.Printf("failed to get a lesson: %v", result.Error)
		return &pb.GetLessonReply{}, result.Error
	}
	return &pb.GetLessonReply{
		Uuid:           lesson.UUID.String(),
		UserId:         lesson.USER_ID.String(),
		CourseId:       lesson.COURSE_ID.String(),
		SequenceNumber: lesson.SEQUENCE_NUMBER,
		Title:          lesson.TITLE,
		Introduction:   lesson.INTRODUCTION,
		Body:           lesson.BODY,
		DefaultCode:    lesson.DEFAULT_CODE,
		AnswerCode:     lesson.ANSWER_CODE,
		Language:       lesson.LANGUAGE,
	}, nil
}

func (s *server) CreateLesson(ctx context.Context, in *pb.CreateLessonRequest) (*pb.CreateLessonReply, error) {
	uUID, _, err := jwt.ParseJWT(in.GetToken())
	if err != nil {

		log.Printf("failed to create lesson: %v", err)
		return &pb.CreateLessonReply{}, err
	}
	courseID, err := uuid.FromString(in.GetCourseId())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}

	lesson := db.Lesson{
		UUID:            uuid.Must(uuid.NewV4()),
		USER_ID:         uUID,
		COURSE_ID:       courseID,
		SEQUENCE_NUMBER: in.GetSequenceNumber(),
		TITLE:           in.GetTitle(),
		INTRODUCTION:    in.GetIntroduction(),
		BODY:            in.GetBody(),
		DEFAULT_CODE:    in.GetDefaultCode(),
		ANSWER_CODE:     in.GetAnswerCode(),
		LANGUAGE:        in.GetLanguage(),
	}
	log.Println(lesson)
	result := db.LessonDB.Create(&lesson)
	if result.Error != nil {
		log.Printf("failed to create lesson: %v", result.Error)
		return &pb.CreateLessonReply{}, result.Error
	}

	return &pb.CreateLessonReply{
		Uuid:           lesson.UUID.String(),
		UserId:         lesson.USER_ID.String(),
		CourseId:       lesson.COURSE_ID.String(),
		SequenceNumber: lesson.SEQUENCE_NUMBER,
		Title:          lesson.TITLE,
		Introduction:   lesson.INTRODUCTION,
		Body:           lesson.BODY,
		DefaultCode:    lesson.DEFAULT_CODE,
		AnswerCode:     lesson.ANSWER_CODE,
		Language:       lesson.LANGUAGE,
	}, nil
}

func (s *server) UpdateLesson(ctx context.Context, in *pb.UpdateLessonRequest) (*pb.UpdateLessonReply, error) {
	var lesson db.Lesson
	userUuid, _, err := jwt.ParseJWT(in.GetToken())
	if err != nil {

		log.Printf("failed to create lesson: %v", err)
		return &pb.UpdateLessonReply{}, err
	}

	uUID, err := uuid.FromString(in.GetUuid())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}
	result := db.LessonDB.First(&lesson, "UUID = ?", uUID)
	if result.Error != nil {
		log.Printf("failed to update lesson: %v", result.Error)
		return &pb.UpdateLessonReply{}, result.Error
	}
	if userUuid != lesson.USER_ID {
		err := errors.New("invalid user id")
		log.Printf("failed to update lesson: %v", err)
		return &pb.UpdateLessonReply{}, err
	}

	lesson.SEQUENCE_NUMBER = in.GetSequenceNumber()
	lesson.TITLE = in.GetTitle()
	lesson.INTRODUCTION = in.GetIntroduction()
	lesson.BODY = in.GetBody()
	lesson.DEFAULT_CODE = in.GetDefaultCode()
	lesson.ANSWER_CODE = in.GetAnswerCode()
	lesson.LANGUAGE = in.GetLanguage()
	db.LessonDB.Save(&lesson)

	return &pb.UpdateLessonReply{
		Uuid:           lesson.UUID.String(),
		UserId:         lesson.USER_ID.String(),
		CourseId:       lesson.COURSE_ID.String(),
		SequenceNumber: lesson.SEQUENCE_NUMBER,
		Title:          lesson.TITLE,
		Introduction:   lesson.INTRODUCTION,
		Body:           lesson.BODY,
		DefaultCode:    lesson.DEFAULT_CODE,
		AnswerCode:     lesson.ANSWER_CODE,
		Language:       lesson.LANGUAGE,
	}, nil
}

func (s *server) DeleteLesson(ctx context.Context, in *pb.DeleteLessonRequest) (*empty.Empty, error) {
	var lesson db.Lesson
	userUuid, _, err := jwt.ParseJWT(in.GetToken())

	if err != nil {

		return new(empty.Empty), err
	}
	uUID, err := uuid.FromString(in.GetUuid())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}
	result := db.LessonDB.First(&lesson, "UUID = ?", uUID)
	if result.Error != nil {

		return new(empty.Empty), result.Error
	}
	if userUuid != lesson.USER_ID {
		return new(empty.Empty), errors.New("invalid user id")
	}

	result = db.LessonDB.Delete(&lesson, "UUID = ?", uUID)
	log.Println(lesson, result.Error)
	if result.Error != nil {
		log.Printf("failed to delete a lesson: %v", result.Error)
		return new(empty.Empty), result.Error
	}

	return new(empty.Empty), nil
}

func RunServer() {
	db.LessonInit()
	defer db.LessonClose()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLessonServer(s, &server{})
	log.Println("lesson grpc server running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
