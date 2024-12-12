package service

import "github.com/drizzleent/nft-transactor/repository"

type TokenService interface {
	CreateToken() error
	ListToken() error
	TotalSupplyToken() error
}

type tokenService struct {
	repository repository.TokenRepository
}

func NewTokenService(r repository.TokenRepository) TokenService {
	return &tokenService{
		repository: r,
	}
}

func (s *tokenService) CreateToken() error {
	return nil
}

func (s *tokenService) ListToken() error {
	return nil
}

func (s *tokenService) TotalSupplyToken() error {
	return nil
}
