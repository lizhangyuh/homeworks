package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "nvm.com/mrc-api-go/api/post/service/v1"
	"nvm.com/mrc-api-go/app/post/service/internal/biz"
)

// PostService is a Post service.
type PostService struct {
	v1.UnimplementedPostServer

	uc  *biz.PostUsecase
	log *log.Helper
}

// NewPostService new a Post service.
func NewPostService(uc *biz.PostUsecase, logger log.Logger) *PostService {
	return &PostService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/post")),
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.CreatePostReply, error) {
	post, err := s.uc.CreatePost(ctx, &biz.Post{
		Title:   req.Title,
		Content: req.Content,
	})

	if err != nil {
		return nil, err
	}

	return &v1.CreatePostReply{
		Id:        post.Id,
		Title:     post.Title,
		Content:   post.Content,
		Published: post.Published,
	}, nil
}

func (s *PostService) PostIndex(ctx context.Context, req *v1.ListPostRequest) (*v1.ListPostReply, error) {
	posts, err := s.uc.ListPost(ctx, req.Title)

	if err != nil {
		return nil, err
	}

	data := make([]*v1.ListPostReply_Post, 0)

	for _, p := range posts {
		data = append(data, &v1.ListPostReply_Post{
			Id:        p.Id,
			Title:     p.Title,
			UserName:  p.UserName,
			Published: p.Published,
		})
	}

	return &v1.ListPostReply{
		Data: data,
	}, nil
}

func (s *PostService) PostShow(ctx context.Context, req *v1.GetPostRequest) (*v1.GetPostReply, error) {
	post, err := s.uc.GetPost(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &v1.GetPostReply{
		Id:        post.Id,
		Title:     post.Title,
		Content:   post.Content,
		Published: post.Published,
	}, nil
}

func (s *PostService) UpdatePost(ctx context.Context, req *v1.UpdatePostRequest) (*v1.UpdatePostReply, error) {
	post, err := s.uc.UpdatePost(ctx, &biz.Post{
		Title:     req.Title,
		Content:   req.Content,
		Published: req.Published,
		UserId:    req.UserId,
	})

	if err != nil {
		return nil, err
	}

	return &v1.UpdatePostReply{
		Id: post.Id,
	}, nil
}
