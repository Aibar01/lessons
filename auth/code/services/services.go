package services

import "auth/code/repository"

type Service struct {
	UserServiceInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserServiceInterface: NewUserService(repo.UserRepositoryInterface),
	}
}
