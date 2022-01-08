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

func getCases(lesson db.Lesson) ([]*pb.Case, error) {
	var pbCases []*pb.Case
	uUID := lesson.UUID
	var dbCases []db.TestCase
	result := db.CaseDB.Find(&dbCases, "LESSON_ID = ?", uUID)
	if result.Error != nil {
		log.Printf("failed to get cases: %v", result.Error)
	}
	for _, c := range dbCases {
		pbCases = append(pbCases, &pb.Case{Input: c.INPUT, Output: c.OUTPUT})
	}
	return pbCases, result.Error
}

func getCasesWithID(lesson db.Lesson) ([]*pb.CaseWithID, error) {
	var pbCases []*pb.CaseWithID
	uUID := lesson.UUID
	var dbCases []db.TestCase
	result := db.CaseDB.Find(&dbCases, "LESSON_ID = ?", uUID)
	if result.Error != nil {
		log.Printf("failed to get cases: %v", result.Error)
	}
	for _, c := range dbCases {
		pbCases = append(pbCases, &pb.CaseWithID{LessonId: c.LESSON_ID.String(), Input: c.INPUT, Output: c.OUTPUT})
	}
	return pbCases, result.Error
}

func (s *server) ListLessons(rect *pb.ListLessonsRequest, stream pb.Lesson_ListLessonsServer) error {
	log.Println("ListLessons running")
	var lessons []db.Lesson
	result := db.LessonDB.Find(&lessons)
	log.Println("got all lessons")
	log.Println(result.Error)
	if result.Error != nil {
		log.Fatalf("failed to list lessons: %v", result.Error)
		return result.Error
	}

	for _, lesson := range lessons {
		pbCases, err := getCases(lesson)
		if err != nil {
			log.Printf("failed to get cases: %v", err)
		}
		if err := stream.Send(&pb.ListLessonsReply{
			Uuid:         lesson.UUID.String(),
			UserId:       lesson.USER_ID.String(),
			CourseId:     lesson.COURSE_ID.String(),
			Title:        lesson.TITLE,
			Introduction: lesson.INTRODUCTION,
			Body:         lesson.BODY,
			DefaultCode:  lesson.DEFAULT_CODE,
			AnswerCode:   lesson.ANSWER_CODE,
			TestCase:     pbCases,
			Language:     lesson.LANGUAGE,
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
	pbCases, err := getCases(lesson)
	if err != nil {
		log.Printf("failed to get cases: %v", err)
	}
	return &pb.GetLessonReply{
		Uuid:         lesson.UUID.String(),
		UserId:       lesson.USER_ID.String(),
		CourseId:     lesson.COURSE_ID.String(),
		Title:        lesson.TITLE,
		Introduction: lesson.INTRODUCTION,
		Body:         lesson.BODY,
		DefaultCode:  lesson.DEFAULT_CODE,
		AnswerCode:   lesson.ANSWER_CODE,
		TestCase:     pbCases,
		Language:     lesson.LANGUAGE,
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
		UUID:         uuid.Must(uuid.NewV4()),
		USER_ID:      uUID,
		COURSE_ID:    courseID,
		TITLE:        in.GetTitle(),
		INTRODUCTION: in.GetIntroduction(),
		BODY:         in.GetBody(),
		DEFAULT_CODE: in.GetDefaultCode(),
		ANSWER_CODE:  in.GetAnswerCode(),
		LANGUAGE:     in.GetLanguage(),
	}
	log.Println(lesson)
	result := db.LessonDB.Create(&lesson)
	if result.Error != nil {
		log.Printf("failed to create lesson: %v", result.Error)
		return &pb.CreateLessonReply{}, result.Error
	}

	log.Printf("in.GetTestCase: %v", in.GetTestCase())
	for _, t := range in.GetTestCase() {
		dbCase := db.TestCase{
			UUID:      uuid.Must(uuid.NewV4()),
			LESSON_ID: lesson.UUID,
			INPUT:     t.Input,
			OUTPUT:    t.Output,
		}
		result := db.CaseDB.Create(&dbCase)
		log.Printf("dbCase: %v", dbCase)
		if result.Error != nil {
			log.Printf("failed to create cases: %v", result.Error)
			return &pb.CreateLessonReply{}, result.Error
		}
	}

	pbCases, err := getCases(lesson)
	if err != nil {
		log.Printf("failed to get cases: %v", err)
	}
	log.Printf("pbCases: %v", pbCases)

	return &pb.CreateLessonReply{
		Uuid:         lesson.UUID.String(),
		UserId:       lesson.USER_ID.String(),
		CourseId:     lesson.COURSE_ID.String(),
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

	lesson.TITLE = in.GetTitle()
	lesson.INTRODUCTION = in.GetIntroduction()
	lesson.BODY = in.GetBody()
	lesson.DEFAULT_CODE = in.GetDefaultCode()
	lesson.ANSWER_CODE = in.GetAnswerCode()
	lesson.LANGUAGE = in.GetLanguage()
	db.LessonDB.Save(&lesson)

	var newDbCases []db.TestCase
	pbCases := in.GetTestCase()
	for _, c := range pbCases {
		newDbCases = append(newDbCases, db.TestCase{INPUT: c.Input, OUTPUT: c.Output})
	}

	var dbCases []db.TestCase
	result = db.CaseDB.First(&dbCases, "LESSON_ID = ?", uUID)
	if result.Error != nil {
		log.Printf("failed to update case: %v", result.Error)
		return &pb.UpdateLessonReply{}, result.Error
	}
	for i, c := range dbCases {
		c.INPUT = newDbCases[i].INPUT
		c.OUTPUT = newDbCases[i].OUTPUT
		db.CaseDB.Save(&c)
	}

	pbCases, err = getCasesWithID(lesson)
	if err != nil {
		log.Printf("failed to get cases: %v", err)
	}

	return &pb.UpdateLessonReply{
		Uuid:         lesson.UUID.String(),
		UserId:       lesson.USER_ID.String(),
		CourseId:     lesson.COURSE_ID.String(),
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
	var cases []db.TestCase
	result = db.CaseDB.Find(&cases, "LESSON_ID = ?", lesson.UUID)
	if result.Error != nil {
		log.Printf("failed to find cases: %v", result.Error)
		return new(empty.Empty), result.Error
	}
	for _, c := range cases {
		result = db.CaseDB.Delete(&c, "UUID = ?", c.UUID)
		if result.Error != nil {
			log.Printf("failed to delete a case: %v", result.Error)
			return new(empty.Empty), result.Error
		}
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
	db.CaseInit()
	defer db.CaseClose()

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
