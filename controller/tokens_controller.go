package controller

import (
	"net/http"

	"github.com/drizzleent/nft-transactor/converter"
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

	req, err := converter.FromRequestToCreateTokenRequest(r)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := tc.service.CreateToken(req)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, resp)

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

	resp, err := tc.service.TotalSupplyToken()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{"result": resp}
	RespondWithJSON(w, http.StatusOK, response)
}
