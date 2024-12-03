package model

type Like struct {
	ID   string `json:"id"`
	User *User  `json:"user"`
	Post *Post  `json:"post"`
}
