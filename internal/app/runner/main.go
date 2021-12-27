package runner

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	pb "github.com/Asuha-a/ProgrammingCourseMarket/internal/pkg/pb/runner"
	"google.golang.org/grpc"
)

const (
	port = ":50055"
)

type server struct {
	pb.UnimplementedRunnerServer
}

type Responce struct {
	Output     string `json:"output"`
	StatusCode int    `json:"statusCode"`
	Memory     string `json:"memory"`
	CpuTime    string `json:"cpuTime"`
}

func (s *server) RunCode(ctx context.Context, in *pb.RunCodeRequest) (*pb.RunCodeReply, error) {
	postBody, _ := json.Marshal(map[string]string{
		"clientId":     os.Getenv("CLIENT_ID"),
		"clientSecret": os.Getenv("CLIENT_SECRET"),
		"script":       in.GetCode(),
		"stdin":        in.GetInput(),
		"language":     in.GetLanguage(),
		"versionIndex": in.GetVersion(),
	})
	responseBody := bytes.NewBuffer(postBody)

	res, err := http.Post("https://api.jdoodle.com/v1/execute", "application/json", responseBody)

	if err != nil {
		log.Printf("failed to request jdoodle: %v", err)
		return &pb.RunCodeReply{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read body: %v", err)
		return &pb.RunCodeReply{}, err
	}
	var responce Responce
	err = json.Unmarshal(body, &responce)
	if err != nil {
		log.Printf("failed to unmarshal: %v", err)
		return &pb.RunCodeReply{}, err
	}
	return &pb.RunCodeReply{
		Output: responce.Output,
	}, nil
}

func RunServer() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRunnerServer(s, &server{})
	log.Println("runner grpc server running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
