package controller

import (
	"net/http"

	"github.com/drizzleent/nft-transactor/service"
)

type TokenController struct {
	service service.TokenService
}

func NewTokenController(srv service.TokenService) *TokenController {
	return &TokenController{
		service: srv,
	}
}

func (tc *TokenController) CreateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}

func (tc *TokenController) ListToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}

func (tc *TokenController) TotalSupplyToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}
