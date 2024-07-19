package grpcsvc

import (
	"context"
	pb "grpc-implementation/grpc-server/pb/story"
)

type StoryService struct {
	// default method from story service protobuf
	// if we don't implement method yet, clients are able to call method stub but return error
	pb.UnimplementedStoryServiceServer
}

func NewStoryService() *StoryService {
	return &StoryService{}
}

func (s *StoryService) FindAll(ctx context.Context, in *pb.FindAllStoriesRequest) (*pb.Stories, error) {
	// Assume this payload is from usecase story
	stories := []*pb.Story{
		{
			Id:      1,
			Title:   "Title 1",
			Content: "Content 1",
		},
		{
			Id:      2,
			Title:   "Title 2",
			Content: "Content 2",
		},
		{
			Id:      3,
			Title:   "Title 3",
			Content: "Content 3",
		},
	}

	return &pb.Stories{
		Stories: stories,
	}, nil
}

func (s *StoryService) FindByID(ctx context.Context, in *pb.FindStoryByIDRequest) (*pb.Story, error) {
	// Assume this payload is from usecase story
	story := &pb.Story{
		Id:      1,
		Title:   "Title 1",
		Content: "Content 1",
	}

	return story, nil
}
