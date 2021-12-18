package lesson

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	result := db.DB.Find(&lessons)
	log.Println("got all lessons")
	log.Println(result.Error)
	if result.Error != nil {
		log.Fatalf("failed to list lessons: %v", result.Error)
		return result.Error
	}

	for _, lesson := range lessons {
		cases := []*pb.Case{}
		for i, t := range lesson.TestCase {
			bytes := []byte(fmt.Sprintf("%v", t))
			err := json.Unmarshal(bytes, &cases[i])
			if err != nil {
				return err
			}
		}
		if err := stream.Send(&pb.ListLessonsReply{
			Uuid:         lesson.UUID.String(),
			UserId:       lesson.USER_ID.String(),
			Title:        lesson.TITLE,
			Introduction: lesson.INTRODUCTION,
			Body:         lesson.BODY,
			DefaultCode:  lesson.DEFAULT_CODE,
			AnswerCode:   lesson.ANSWER_CODE,
			TestCase:     cases,
			Language:     lesson.LANGUAGE,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetLesson(ctx context.Context, in *pb.GetLessonRequest) (*pb.GetLessonReply, error) {
	var lesson db.Lesson
	result := db.DB.First(&lesson, "UUID = ?", in.GetUuid())
	log.Println(lesson)
	if result.Error != nil {
		log.Printf("failed to get a lesson: %v", result.Error)
		return &pb.GetLessonReply{}, result.Error
	}
	cases := []*pb.Case{}
	for i, t := range lesson.TestCase {
		bytes := []byte(fmt.Sprintf("%v", t))
		err := json.Unmarshal(bytes, &cases[i])
		if err != nil {
			return &pb.GetLessonReply{}, err
		}
	}
	return &pb.GetLessonReply{
		Uuid:         lesson.UUID.String(),
		UserId:       lesson.USER_ID.String(),
		Title:        lesson.TITLE,
		Introduction: lesson.INTRODUCTION,
		Body:         lesson.BODY,
		DefaultCode:  lesson.DEFAULT_CODE,
		AnswerCode:   lesson.ANSWER_CODE,
		TestCase:     cases,
		Language:     lesson.LANGUAGE,
	}, nil
}

func (s *server) CreateLesson(ctx context.Context, in *pb.CreateLessonRequest) (*pb.CreateLessonReply, error) {
	uUID, _, err := jwt.ParseJWT(in.GetToken())
	if err != nil {
		log.Println(err)
		log.Printf("failed to create lesson: %v", err)
		return &pb.CreateLessonReply{}, err
	}
	courseID, err := uuid.FromString(in.GetCourseId())
	dbCases := []db.TestCase{}
	log.Printf("in.GetTestCase: %v", in.GetTestCase())
	for _, t := range in.GetTestCase() {
		bytes := []byte(fmt.Sprintf("%v", t))
		var dbCase db.TestCase
		err := json.Unmarshal(bytes, &dbCase)
		if err != nil {
			log.Printf("failed to unmarshal[db]: %v", err)
			return &pb.CreateLessonReply{}, err
		}
		dbCases = append(dbCases, dbCase)
	}
	lesson := db.Lesson{
		UUID:         uuid.Must(uuid.NewV4()),
		USER_ID:      uUID,
		COURSE_ID:    courseID,
		TITLE:        in.GetTitle(),
		INTRODUCTION: in.GetIntroduction(),
		BODY:         in.GetBody(),
		DEFAULT_CODE: in.GetDefaultCode(),
		ANSWER_CODE:  in.GetAnswerCode(),
		TestCase:     dbCases,
		LANGUAGE:     in.GetLanguage(),
	}
	log.Println(lesson)
	result := db.DB.Create(&lesson)
	if result.Error != nil {
		log.Printf("failed to create lesson: %v", result.Error)
		return &pb.CreateLessonReply{}, result.Error
	}

	pbCases := []*pb.Case{}
	for _, t := range lesson.TestCase {
		bytes := []byte(fmt.Sprintf("%v", t))
		var pbCase *pb.Case
		err := json.Unmarshal(bytes, &pbCase)
		if err != nil {
			log.Printf("failed to unmarshal[pb]: %v", err)
			return &pb.CreateLessonReply{}, err
		}
		pbCases = append(pbCases, pbCase)
	}

	return &pb.CreateLessonReply{
		Uuid:         lesson.UUID.String(),
		UserId:       lesson.USER_ID.String(),
		Title:        lesson.TITLE,
		Introduction: lesson.INTRODUCTION,
		Body:         lesson.BODY,
		DefaultCode:  lesson.DEFAULT_CODE,
		AnswerCode:   lesson.ANSWER_CODE,
		TestCase:     pbCases,
		Language:     lesson.LANGUAGE,
	}, nil
}

func (s *server) UpdateLesson(ctx context.Context, in *pb.UpdateLessonRequest) (*pb.UpdateLessonReply, error) {
	var lesson db.Lesson
	userUuid, _, err := jwt.ParseJWT(in.GetToken())
	if err != nil {
		log.Println(err)
		log.Printf("failed to create lesson: %v", err)
		return &pb.UpdateLessonReply{}, err
	}

	uUID, err := uuid.FromString(in.GetUuid())
	result := db.DB.First(&lesson, "UUID = ?", uUID)
	if result.Error != nil {
		log.Printf("failed to update lesson: %v", result.Error)
		return &pb.UpdateLessonReply{}, result.Error
	}
	if userUuid != lesson.USER_ID {
		err := errors.New("invalid user id")
		log.Printf("failed to update lesson: %v", err)
		return &pb.UpdateLessonReply{}, err
	}
	dbCases := []db.TestCase{}
	for i, t := range in.GetTestCase() {
		bytes := []byte(fmt.Sprintf("%v", t))
		err := json.Unmarshal(bytes, &dbCases[i])
		if err != nil {
			return &pb.UpdateLessonReply{}, err
		}
	}
	lesson.TITLE = in.GetTitle()
	lesson.INTRODUCTION = in.GetIntroduction()
	lesson.BODY = in.GetBody()
	lesson.DEFAULT_CODE = in.GetDefaultCode()
	lesson.ANSWER_CODE = in.GetAnswerCode()
	lesson.TestCase = dbCases
	lesson.LANGUAGE = in.GetLanguage()
	db.DB.Save(&lesson)

	pbCases := []*pb.Case{}
	for i, t := range lesson.TestCase {
		bytes := []byte(fmt.Sprintf("%v", t))
		err := json.Unmarshal(bytes, &pbCases[i])
		if err != nil {
			return &pb.UpdateLessonReply{}, err
		}
	}

	return &pb.UpdateLessonReply{
		Uuid:         lesson.UUID.String(),
		UserId:       lesson.USER_ID.String(),
		Title:        lesson.TITLE,
		Introduction: lesson.INTRODUCTION,
		Body:         lesson.BODY,
		DefaultCode:  lesson.DEFAULT_CODE,
		AnswerCode:   lesson.ANSWER_CODE,
		TestCase:     pbCases,
		Language:     lesson.LANGUAGE,
	}, nil
}

func (s *server) DeleteLesson(ctx context.Context, in *pb.DeleteLessonRequest) (*empty.Empty, error) {
	var lesson db.Lesson
	userUuid, _, err := jwt.ParseJWT(in.GetToken())

	if err != nil {
		log.Println(err)
		return new(empty.Empty), err
	}
	uUID, err := uuid.FromString(in.GetUuid())
	result := db.DB.First(&lesson, "UUID = ?", uUID)
	if result.Error != nil {
		log.Println(err)
		return new(empty.Empty), result.Error
	}
	if userUuid != lesson.USER_ID {
		return new(empty.Empty), errors.New("invalid user id")
	}
	result = db.DB.Delete(&lesson, "UUID = ?", uUID)
	log.Println(lesson, result.Error)
	if result.Error != nil {
		log.Printf("failed to delete a lesson: %v", result.Error)
		return new(empty.Empty), result.Error
	}
	return new(empty.Empty), nil
}

func RunServer() {
	log.Println("test")
	db.Init()
	defer db.Close()

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
