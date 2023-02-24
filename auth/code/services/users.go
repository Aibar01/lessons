package services

import (
	"auth/code/domains"
	"auth/code/repository"
	"context"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
	CreateUser(context context.Context, data *domains.User) error
	GetUser(context context.Context, id uuid.UUID) (*domains.User, error)
}

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(context context.Context, data *domains.User) error {
	return u.repo.CreateUser(context, data)
}

func (u *UserService) GetUser(context context.Context, id uuid.UUID) (*domains.User, error) {
	return u.repo.GetUser(context, id)
}
