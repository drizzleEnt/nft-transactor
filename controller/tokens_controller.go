package controller

import "net/http"

type TokenController struct {
}

func NewTokenController() *TokenController {
	return &TokenController{}
}

func (tc *TokenController) CreateToken(w http.ResponseWriter, r *http.Request) {

}

func (tc *TokenController) ListToken(w http.ResponseWriter, r *http.Request) {

}

func (tc *TokenController) TotalSupplyToken(w http.ResponseWriter, r *http.Request) {

}
