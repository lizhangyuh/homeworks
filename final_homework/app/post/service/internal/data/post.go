package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"nvm.com/mrc-api-go/app/post/service/internal/biz"
)

type PostRepo struct {
	data *Data
	log  *log.Helper
}

// NewPostRepo .
func NewPostRepo(data *Data, logger log.Logger) biz.PostRepo {
	return &PostRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *PostRepo) Create(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	res := r.data.db.Create(&post)

	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}

	return post, nil
}

func (r *PostRepo) Update(ctx context.Context, post *biz.Post) (*biz.Post, error) {
	res := r.data.db.
		Update("title", post.Title).
		Update("content", post.Content).
		Update("published", post.Published)

	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}

	return post, nil
}

func (r *PostRepo) FindByID(cxt context.Context, id int64) (*biz.Post, error) {
	var post biz.Post
	result := r.data.db.Where(&biz.Post{Id: id}).First(&post)

	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}

func (r *PostRepo) ListByTitle(ctx context.Context, title string) ([]*biz.Post, error) {
	var posts []*biz.Post
	results := r.data.db.Where(&biz.Post{Title: title}).Find(posts)

	if results.Error != nil {
		return nil, results.Error
	}

	return posts, nil
}

func (r *PostRepo) ListAll(context.Context) ([]*biz.Post, error) {
	var posts []*biz.Post
	results := r.data.db.Find(posts)

	if results.Error != nil {
		return nil, results.Error
	}

	return posts, nil
}
