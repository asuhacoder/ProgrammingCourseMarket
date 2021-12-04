package main

import (
	"context"
	"log"
	"net"

	db "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/db/user"
	pb "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/auth"
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct {
	pb.UnimplementedAuthServer
}

type userClaims struct {
	UUID       uuid.UUID
	PERMISSION string
	jwt.StandardClaims
}

func createJWT(user db.User) (string, error) {
	mySingningKey := []byte("AllYourBase")

	claims := userClaims{
		user.UUID,
		user.PERMISSION,
		jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySingningKey)

	return ss, err
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	var user db.User
	result := db.DB.Where("email = ?", in.GetEmail()).First(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(in.GetPassword()))
	if err != nil {
		panic(err)
	}

	ss, err := createJWT(user)
	if err != nil {
		panic(err)
	}

	return &pb.LoginReply{Token: ss}, nil
}

func main() {
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
