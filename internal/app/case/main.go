package lesson

import (
	"context"
	"errors"
	"log"
	"net"

	db "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/db/case"
	"github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/jwt"
	pb "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/case"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const (
	port = ":50056"
)

type server struct {
	pb.UnimplementedCaseServer
}

func (s *server) ListCases(rect *pb.ListCasesRequest, stream pb.Case_ListCasesServer) error {
	log.Println("ListCases running")
	var testCases []db.TestCase
	result := db.CaseDB.Where("Lesson_ID = ?", rect.GetLessonId()).Find(&testCases)
	log.Println("got all testCases")
	log.Println(result.Error)
	if result.Error != nil {
		log.Fatalf("failed to list testCases: %v", result.Error)
		return result.Error
	}

	for _, testCase := range testCases {
		if err := stream.Send(&pb.ListCasesReply{
			Uuid:     testCase.UUID.String(),
			UserId:   testCase.USER_ID.String(),
			LessonId: testCase.LESSON_ID.String(),
			Input:    testCase.INPUT,
			Output:   testCase.OUTPUT,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetCase(ctx context.Context, in *pb.GetCaseRequest) (*pb.GetCaseReply, error) {
	var testCase db.TestCase
	result := db.CaseDB.Where("UUID = ?", in.GetUuid()).First(&testCase)
	log.Println(testCase)
	if result.Error != nil {
		log.Printf("failed to get a case: %v", result.Error)
		return &pb.GetCaseReply{}, result.Error
	}
	return &pb.GetCaseReply{
		Uuid:     testCase.UUID.String(),
		UserId:   testCase.USER_ID.String(),
		LessonId: testCase.LESSON_ID.String(),
		Input:    testCase.INPUT,
		Output:   testCase.OUTPUT,
	}, nil
}

func (s *server) CreateCase(ctx context.Context, in *pb.CreateCaseRequest) (*pb.CreateCaseReply, error) {
	uUID, _, err := jwt.ParseJWT(in.GetToken())
	if err != nil {

		log.Printf("failed to create case: %v", err)
		return &pb.CreateCaseReply{}, err
	}
	lessonID, err := uuid.FromString(in.GetLessonId())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}
	dbCase := db.TestCase{
		UUID:      uuid.Must(uuid.NewV4()),
		USER_ID:   uUID,
		LESSON_ID: lessonID,
		INPUT:     in.GetInput(),
		OUTPUT:    in.GetOutput(),
	}
	result := db.CaseDB.Create(&dbCase)
	log.Printf("dbCase: %v", dbCase)
	if result.Error != nil {
		log.Printf("failed to create cases: %v", result.Error)
		return &pb.CreateCaseReply{}, result.Error
	}

	return &pb.CreateCaseReply{
		Uuid:     dbCase.UUID.String(),
		UserId:   dbCase.USER_ID.String(),
		LessonId: dbCase.LESSON_ID.String(),
		Input:    dbCase.INPUT,
		Output:   dbCase.OUTPUT,
	}, nil
}

func (s *server) UpdateCase(ctx context.Context, in *pb.UpdateCaseRequest) (*pb.UpdateCaseReply, error) {
	var testCase db.TestCase
	userUuid, _, err := jwt.ParseJWT(in.GetToken())
	if err != nil {

		log.Printf("failed to create case: %v", err)
		return &pb.UpdateCaseReply{}, err
	}

	uUID, err := uuid.FromString(in.GetUuid())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}
	result := db.CaseDB.First(&testCase, "UUID = ?", uUID)
	if result.Error != nil {
		log.Printf("failed to update case: %v", result.Error)
		return &pb.UpdateCaseReply{}, result.Error
	}
	if userUuid != testCase.USER_ID {
		err := errors.New("invalid user id")
		log.Printf("failed to update case: %v", err)
		return &pb.UpdateCaseReply{}, err
	}

	testCase.INPUT = in.GetInput()
	testCase.OUTPUT = in.GetOutput()
	db.CaseDB.Save(&testCase)

	return &pb.UpdateCaseReply{
		Uuid:     testCase.UUID.String(),
		UserId:   testCase.USER_ID.String(),
		LessonId: testCase.LESSON_ID.String(),
		Input:    testCase.INPUT,
		Output:   testCase.OUTPUT,
	}, nil
}

func (s *server) DeleteCase(ctx context.Context, in *pb.DeleteCaseRequest) (*empty.Empty, error) {
	var testCase db.TestCase
	userUuid, _, err := jwt.ParseJWT(in.GetToken())

	if err != nil {

		return new(empty.Empty), err
	}
	uUID, err := uuid.FromString(in.GetUuid())
	if err != nil {
		log.Printf("failed to convert string to uuid: %v", err)
	}
	result := db.CaseDB.First(&testCase, "UUID = ?", uUID)
	if result.Error != nil {
		return new(empty.Empty), result.Error
	}
	if userUuid != testCase.USER_ID {
		return new(empty.Empty), errors.New("invalid user id")
	}

	result = db.CaseDB.Delete(&testCase, "UUID = ?", uUID)
	log.Println(testCase, result.Error)
	if result.Error != nil {
		log.Printf("failed to delete a case %v", result.Error)
		return new(empty.Empty), result.Error
	}

	return new(empty.Empty), nil
}

func RunServer() {
	db.CaseInit()
	defer db.CaseClose()
	db.CaseInit()
	defer db.CaseClose()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCaseServer(s, &server{})
	log.Println("case grpc server running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
