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

func (s *server) Authn(ctx context.Context, in *pb.AuthnRequest) (*pb.AuthnReply, error) {
	var user db.User
	result := db.DB.Where("email = ?", in.GetEmail()).First(&user)
	if result.Error != nil {
		log.Printf("failed to get user: %v", result.Error)
		return &pb.AuthnReply{}, result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(in.GetPassword()))
	if err != nil {
		log.Printf("failed to login: %v", err)
	}

	ss, err := jwt.CreateJWT(user)
	if err != nil {
		panic(err)
	}
	return &pb.AuthnReply{
		Token:      ss,
		Uuid:       user.UUID.String(),
		Email:      user.EMAIL,
		Permission: user.PERMISSION,
	}, nil
}

func (s *server) Authz(ctx context.Context, in *pb.AuthzRequest) (*pb.AuthzReply, error) {
	var user db.User
	uuid, _, err := jwt.ParseJWT(in.GetToken())
	if err != nil {
		log.Printf("invalid token: %v", err)
	}
	result := db.DB.Where("uuid = ?", uuid).First(&user)
	if result.Error != nil {
		log.Printf("failed to get user: %v", result.Error)
		return &pb.AuthzReply{}, result.Error
	}

	return &pb.AuthzReply{
		Token:      in.GetToken(),
		Uuid:       user.UUID.String(),
		Email:      user.EMAIL,
		Permission: user.PERMISSION,
	}, nil
}

func RunServer() {
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
