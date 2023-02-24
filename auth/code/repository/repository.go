package repository

import (
	"auth/code/domains"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	UserRepositoryInterface
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UserRepositoryInterface: NewUserRepository(db),
	}
}

func createDomain(context context.Context, data domains.Domain, db *pgxpool.Pool) error {
	result := make([][]any, 1)
	result[0] = data.GetDBValues()

	_, err := db.CopyFrom(
		context,
		pgx.Identifier{data.GetDBTableName()},
		domains.GetDomainFields(data),
		pgx.CopyFromRows(result),
	)

	if err != nil {
		return err
	}

	return nil
}
