package service

import (
	"context"

	pb "nvm.com/mrc-api-go/api/bff/mobile/v1"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/biz"
)

func (s *MoibleService) PostIndex(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
	posts, err := s.pc.PostIndex(ctx, req.Title)

	if err != nil {
		return nil, err
	}

	data := make([]*pb.ListPostReply_Post, 0)

	for _, p := range posts {
		data = append(data, &pb.ListPostReply_Post{
			Id:        p.Id,
			Title:     p.Title,
			UserName:  p.UserName,
			Published: p.Published,
		})
	}

	return &pb.ListPostReply{
		Data: data,
	}, nil
}

func (s *MoibleService) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	post, err := s.pc.CreatePost(ctx, &biz.Post{
		Title:   req.Title,
		Content: req.Content,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreatePostReply{
		Id:        post.Id,
		Title:     post.Title,
		Content:   post.Content,
		Published: post.Published,
	}, nil
}

func (s *MoibleService) PostShow(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostReply, error) {
	post, err := s.pc.PostShow(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.GetPostReply{
		Id:        post.Id,
		Title:     post.Title,
		Content:   post.Content,
		Published: post.Published,
	}, nil
}

func (s *MoibleService) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostReply, error) {
	post, err := s.pc.UpdatePost(ctx, &biz.Post{
		Title:     req.Title,
		Content:   req.Content,
		Published: req.Published,
		UserId:    req.UserId,
	})

	if err != nil {
		return nil, err
	}

	return &pb.UpdatePostReply{
		Id: post.Id,
	}, nil
}
