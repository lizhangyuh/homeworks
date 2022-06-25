package data

import (
	"context"

	postService "nvm.com/mrc-api-go/api/post/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"nvm.com/mrc-api-go/app/bff/mobile/internal/biz"
)

type postRepo struct {
	data *Data
	log  *log.Helper
}

// NewpostRepo .
func NewpostRepo(data *Data, logger log.Logger) biz.PostRepo {
	return &postRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/post")),
	}
}

func (u *postRepo) CreatePost(c context.Context, post *biz.Post) (*biz.Post, error) {
	createPost, err := u.data.pc.CreatePost(c, &postService.CreatePostRequest{
		Title:     post.Title,
		Content:   post.Content,
		UserId:    post.UserId,
		Published: post.Published,
	})

	if err != nil {
		return nil, err
	}

	return &biz.Post{
		Id:        createPost.Id,
		Title:     createPost.Title,
		Content:   createPost.Content,
		Published: createPost.Published,
	}, nil
}

func (u *postRepo) PostIndex(c context.Context, title string) ([]*biz.Post, error) {
	ps, err := u.data.pc.ListPost(c, &postService.ListPostRequest{
		Title: title,
	})

	if err != nil {
		return nil, err
	}

	posts := make([]*biz.Post, 0)

	for _, p := range ps.Data {
		posts = append(posts, &biz.Post{
			Id:        p.Id,
			Title:     p.Title,
			UserName:  p.UserName,
			Published: p.Published,
		})
	}

	return posts, nil
}

func (u *postRepo) PostShow(c context.Context, id int64) (*biz.Post, error) {
	post, err := u.data.pc.GetPost(c, &postService.GetPostRequest{Id: id})

	if err != nil {
		return nil, err
	}

	return &biz.Post{
		Id:        post.Id,
		Title:     post.Title,
		Content:   post.Content,
		UserId:    post.UserId,
		Published: post.Published,
	}, nil
}

func (u *postRepo) UpdatePost(cxt context.Context, p *biz.Post) (*biz.Post, error) {
	post, err := u.data.pc.UpdatePost(cxt, &postService.UpdatePostRequest{
		Title:     p.Title,
		Content:   p.Content,
		UserId:    p.UserId,
		Published: p.Published,
	})

	if err != nil {
		return nil, err
	}

	return &biz.Post{
		Id: post.Id,
	}, nil
}
