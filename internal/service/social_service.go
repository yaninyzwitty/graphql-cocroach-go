package service

import (
	"context"

	"github.com/yaninyzwitty/graphql-cocroach-go/graph/model"
)

type SocialService interface {
	CreateUser(ctx context.Context, input model.NewUser) (*model.User, error)
	UpdateUser(ctx context.Context, input model.UpdateUser, id string) (*model.User, error)
	DeleteUser(ctx context.Context, id string) (bool, error)
	DeletePost(ctx context.Context, id string) (bool, error)
	CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error)
	UpdatePost(ctx context.Context, input model.UpdatePost, id string) (*model.Post, error)
	CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error)
	UpdateComment(ctx context.Context, input model.UpdateComment, id string) (*model.Comment, error)
	DeleteComment(ctx context.Context, id string) (bool, error)
	LikePost(ctx context.Context, input model.NewLike) (*model.Like, error)
	UnlikePost(ctx context.Context, id string) (bool, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
	GetUsers(ctx context.Context) ([]*model.User, error)
}

type socialService struct {
}

func NewSocialService() SocialService {
	return &socialService{}
}

func (s *socialService) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return &model.User{}, nil
}

func (s *socialService) UpdateUser(ctx context.Context, input model.UpdateUser, id string) (*model.User, error) {
	return &model.User{}, nil

}

func (c *socialService) DeleteUser(ctx context.Context, id string) (bool, error) {
	return false, nil
}
func (c *socialService) DeletePost(ctx context.Context, id string) (bool, error) {
	return false, nil
}

func (c *socialService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return &model.User{}, nil
}
func (c *socialService) GetUsers(ctx context.Context) ([]*model.User, error) {
	return []*model.User{}, nil
}
func (c *socialService) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	return &model.Post{}, nil
}
func (c *socialService) UpdatePost(ctx context.Context, input model.UpdatePost, id string) (*model.Post, error) {
	return &model.Post{}, nil
}

func (c *socialService) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	return &model.Comment{}, nil
}

func (c *socialService) UpdateComment(ctx context.Context, input model.UpdateComment, id string) (*model.Comment, error) {
	return &model.Comment{}, nil
}

func (c *socialService) DeleteComment(ctx context.Context, id string) (bool, error) {
	return false, nil
}

func (c *socialService) LikePost(ctx context.Context, input model.NewLike) (*model.Like, error) {
	return &model.Like{}, nil
}
func (c *socialService) UnlikePost(ctx context.Context, id string) (bool, error) {
	return false, nil
}
