package main

import (
	"context"
	"log"
	"net"

	pb "github.com/GHjiejie/grpc-test/user" // 替换为实际路径
	"google.golang.org/grpc"
)

// server 结构体实现了 UserServiceServer 接口
type server struct {
	pb.UnimplementedUserServiceServer
}

// GetUser 实现 GetUser 方法
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// 示例：根据请求 ID 返回用户信息
	return &pb.GetUserResponse{
		Id:   req.Id,
		Name: "Dummy User", // 这里可以替换为真实逻辑
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
