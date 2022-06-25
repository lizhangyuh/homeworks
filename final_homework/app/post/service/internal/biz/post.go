package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	v1 "nvm.com/mrc-api-go/api/post/service/v1"
)

var (
	ErrAuthFailed   = errors.Unauthorized("auth failed", "用户认证出错")
	ErrPostNotFound = errors.NotFound(v1.PostServiceErrorReason_POST_NOT_FOUND.String(), "没有找到相关文章")
)

type Post struct {
	Id        int64
	Title     string
	UserId    int64
	UserName  string
	Content   string
	Published bool
}

// PostRepo is a Greater repo.
type PostRepo interface {
	Create(context.Context, *Post) (*Post, error)
	Update(context.Context, *Post) (*Post, error)
	FindByID(context.Context, int64) (*Post, error)
	ListByTitle(context.Context, string) ([]*Post, error)
	ListAll(context.Context) ([]*Post, error)
}

// PostUsecase is a Post usecase.
type PostUsecase struct {
	repo PostRepo
	log  *log.Helper
}

// NewPostUsecase new a Post usecase.
func NewPostUsecase(repo PostRepo, logger log.Logger) *PostUsecase {
	return &PostUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PostUsecase) CreatePost(ctx context.Context, post *Post) (*Post, error) {
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwtv4.MapClaims)
		if c["user_id"] == nil {
			return nil, ErrAuthFailed
		}

		userId = int64(c["user_id"].(float64))
	}

	// create post
	post.UserId = userId

	post, err := uc.repo.Create(ctx, post)

	// push to kafka

	return post, err
}

func (uc *PostUsecase) ListPost(ctx context.Context, title string) ([]*Post, error) {
	if len(title) == 0 {
		return uc.repo.ListAll(ctx)
	} else {
		return uc.repo.ListByTitle(ctx, title)
	}
}

func (uc *PostUsecase) GetPost(ctx context.Context, id int64) (*Post, error) {
	post, err := uc.repo.FindByID(ctx, id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrPostNotFound
	}

	return post, nil
}

func (uc *PostUsecase) UpdatePost(ctx context.Context, post *Post) (*Post, error) {
	return uc.repo.Update(ctx, post)
}
