package repository

import (
	"context"
	"database/sql"

	"github.com/drizzleent/nft-transactor/models"
)

type TokenRepository interface {
	CreateToken(context.Context, *models.Token) error
	ListToken(context.Context, *models.RequestParams) ([]models.Token, error)
}

type tokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) TokenRepository {
	return &tokenRepository{
		db: db,
	}
}

func (r *tokenRepository) CreateToken(ctx context.Context, token *models.Token) error {
	query := `INSERT INTO tokens 
	(unique_hash, tx_hash, media_url, owner) VALUES 
	($1, $2, $3, $4)`

	args := []interface{}{token.UniqueHash, token.TxHash, token.MediaUrl, token.Owner}

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *tokenRepository) ListToken(ctx context.Context, params *models.RequestParams) ([]models.Token, error) {
	query := `SELECT unique_hash, tx_hash, media_url, owner FROM tokens LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tokens := make([]models.Token, 0)
	for rows.Next() {
		var token models.Token
		err := rows.Scan(&token.UniqueHash, &token.TxHash, &token.MediaUrl, &token.Owner)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}
	return tokens, nil
}
