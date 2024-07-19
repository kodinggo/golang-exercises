package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "grpc-implementation/grpc-server/pb/story"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const storygRPCHost = "localhost:8081"

func main() {
	// connect to grpc server without credentials
	conn, err := grpc.NewClient(storygRPCHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("failed to open connection grpc server, error %v", err)
	}
	defer conn.Close()

	// init grpc client as package dependency from grpc-server repository
	storyClient := pb.NewStoryServiceClient(conn)

	// call grpc method
	result, err := storyClient.FindAll(context.Background(), &pb.FindAllStoriesRequest{})
	if err != nil {
		log.Fatalf("failed to find all stories, error %v", err)
	}

	bStories, _ := json.Marshal(result.Stories)
	fmt.Println(string(bStories))
}
