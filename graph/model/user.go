package model

import "github.com/google/uuid"

type User struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Posts    []*Post    `json:"posts"`
	Comments []*Comment `json:"comments"`
}
type UserUpdated struct {
	ID       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Posts    []*Post    `json:"posts"`
	Comments []*Comment `json:"comments"`
}
