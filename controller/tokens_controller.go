package controller

import (
	"net/http"

	"github.com/drizzleent/nft-transactor/app"
)

type TokenController struct {
}

func NewTokenController() *TokenController {
	return &TokenController{}
}

func (tc *TokenController) CreateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}

func (tc *TokenController) ListToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}

func (tc *TokenController) TotalSupplyToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}
