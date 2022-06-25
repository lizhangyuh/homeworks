package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"nvm.com/mrc-api-go/pkg/auth"
)

type Post struct {
	Id        int64
	Title     string
	Content   string
	UserName  string
	UserId    int64
	Published bool
}

type PostRepo interface {
	CreatePost(context.Context, *Post) (*Post, error)
	PostIndex(ctx context.Context, title string) ([]*Post, error)
	PostShow(ctx context.Context, id int64) (*Post, error)
	UpdatePost(context.Context, *Post) (*Post, error)
}

type PostUsecase struct {
	repo PostRepo
	log  *log.Helper
}

func NewPostUsecase(repo PostRepo, logger log.Logger) *PostUsecase {
	return &PostUsecase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/client"))}
}

func (uc *PostUsecase) CreatePost(ctx context.Context, post *Post) (*Post, error) {
	userId, err := auth.CurrentUserId(ctx)

	if err != nil {
		return nil, err
	}

	post.UserId = userId.(int64)

	return uc.repo.CreatePost(ctx, post)
}

func (uc *PostUsecase) PostIndex(ctx context.Context, title string) ([]*Post, error) {
	return uc.repo.PostIndex(ctx, title)
}

func (uc *PostUsecase) PostShow(ctx context.Context, id int64) (*Post, error) {
	return uc.repo.PostShow(ctx, id)
}

func (uc *PostUsecase) UpdatePost(ctx context.Context, post *Post) (*Post, error) {
	userId, err := auth.CurrentUserId(ctx)

	if err != nil {
		return nil, err
	}

	post.UserId = userId.(int64)

	return uc.repo.UpdatePost(ctx, post)
}
