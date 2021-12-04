package auth

import (
	"context"
	"log"
	"net"

	db "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/db/user"
	"github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/jwt"
	pb "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/auth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct {
	pb.UnimplementedAuthServer
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	var user db.User
	result := db.DB.Where("email = ?", in.GetEmail()).First(&user)
	if result.Error != nil {
		log.Printf("failed to get user: %v", result.Error)
		return &pb.LoginReply{Token: ""}, result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(in.GetPassword()))
	if err != nil {
		log.Printf("failed to login: %v", err)
	}

	ss, err := jwt.CreateJWT(user)
	if err != nil {
		panic(err)
	}

	return &pb.LoginReply{Token: ss}, nil
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
	pb.RegisterAuthServer(s, &server{})
	log.Println("auth grpc server running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
