package user

import (
	"context"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ID          uint64 `json:"id"`
	Username    string `json:"username"`
}

type Repository interface {
	Login(ctx context.Context, user *User) (*User, error)
	GetByUsername(ctx context.Context, username string)
	Register(ctx context.Context, User *User) (*User, error)
}

type Service interface {
	CreateNewUser(ctx context.Context, req *UserRequest) (*User, error)
	Login(ctx context.Context, req *UserRequest) (*LoginResponse, error)
}
