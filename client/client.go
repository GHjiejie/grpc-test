package main

import (
	"context"
	"log"
	"time"

	pb "github.com/GHjiejie/grpc-test/user" // 替换为实际路径
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 创建请求
	req := &pb.GetUserRequest{Id: "123"}

	// 调用 GetUser 方法
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	log.Printf("User ID: %s, Name: %s", res.Id, res.Name)
}
