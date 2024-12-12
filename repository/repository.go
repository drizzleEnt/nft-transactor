package repository

import "database/sql"

type TokenRepository interface {
	CreateToken() error
	ListToken() error
	TotalSupplyToken() error
}

type tokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) TokenRepository {
	return &tokenRepository{
		db: db,
	}
}

func (r *tokenRepository) CreateToken() error {
	return nil
}

func (r *tokenRepository) ListToken() error {
	return nil
}

func (r *tokenRepository) TotalSupplyToken() error {
	return nil
}
