package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yaninyzwitty/graphql-cocroach-go/graph/model"
)

var (
	name  string
	email string
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
	DB *pgxpool.Pool
}

func NewSocialService(DB *pgxpool.Pool) SocialService {
	return &socialService{
		DB: DB,
	}
}

func (s *socialService) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	if input.Name == "" || input.Email == "" {
		return nil, fmt.Errorf("name and email are required")
	}

	userID := uuid.New().String()

	query := `INSERT INTO users(id, name, email) Values($1, $2, $3)`

	_, err := s.DB.Exec(ctx, query, userID, input.Name, input.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user to cocroach: %w", err)
	}
	createdUser := &model.User{
		ID:    userID,
		Name:  input.Name,
		Email: input.Email,
	}
	return createdUser, nil
}

func (s *socialService) UpdateUser(ctx context.Context, input model.UpdateUser, id string) (*model.User, error) {
	if input.Name != nil {
		name = *input.Name
	}
	if input.Email != nil {
		email = *input.Email
	}
	var updatedUser model.User
	query := `UPDATE users SET (name, email) = ($1, $2) WHERE id = $3 returning id, name, email`

	err := s.DB.QueryRow(ctx, query, name, email, id).Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &updatedUser, nil

}

func (c *socialService) DeleteUser(ctx context.Context, id string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("id is required")

	}

	query := `DELETE FROM users WHERE id = $1`
	_, err := c.DB.Exec(ctx, query, id)
	if err != nil {
		return false, fmt.Errorf("failed to delete user: %w", err)
	}

	return true, nil
}
func (c *socialService) DeletePost(ctx context.Context, id string) (bool, error) {
	return false, nil
}

func (c *socialService) GetUser(ctx context.Context, id string) (*model.User, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	var user model.User
	query := `SELECT id, name, email FROM users WHERE id = $1`

	row := c.DB.QueryRow(ctx, query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user with id: %s not found", id)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}
func (c *socialService) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	query := `SELECT id, name, email FROM users`
	rows, err := c.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := &model.User{}

		// Ensure that ID is a string
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		users = append(users, user)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over users: %w", err)
	}

	return users, nil
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
