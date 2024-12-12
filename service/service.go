package service

type TokenService interface {
}

type tokenService struct {
}

func NewTokenService() TokenService {
	return &tokenService{}
}

func (s *tokenService) CreateToken() {

}

func (s *tokenService) ListToken() {

}

func (s *tokenService) TotalSupplyToken() {

}
