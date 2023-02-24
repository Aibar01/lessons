package repository

import (
	"auth/code/domains"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepositoryInterface interface {
	CreateUser(context context.Context, data *domains.User) error
	GetUser(context context.Context, id uuid.UUID) (*domains.User, error)
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(context context.Context, data *domains.User) error {
	return createDomain(context, data, u.db)
}

func (u *UserRepository) GetUser(context context.Context, id uuid.UUID) (*domains.User, error) {
	//TODO implement me
	panic("implement me")
}
