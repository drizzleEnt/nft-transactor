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

// CreateToken создает токен и записывает параметры в БД
// @Summary Создание нового токена
// @Description Создает уникальный токен в блокчейне и записывает параметры в базу данных
// @Tags tokens
// @Accept json
// @Produce json
// @Param token body models.CreateTokenRequest true "Параметры токена"
// @Success 200 {object} models.Token
// @Failure 400 {object} map[string]string
// @Failure 405 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /create [post]
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

	resp, err := tc.service.CreateToken(r.Context(), req)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, resp)

}

// ListToken возвращает список всех токенов
// @Summary Получение списка токенов
// @Description Возвращает список всех объектов модели Token
// @Tags tokens
// @Accept json
// @Produce json
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} []models.Token
// @Failure 405 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /list [get]
func (tc *TokenController) ListToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	params := converter.FromRequestParamsToModel(r)

	resp, err := tc.service.ListToken(r.Context(), params)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := map[string]interface{}{"Tokens": resp}
	RespondWithJSON(w, http.StatusOK, response)
}

// TotalSupplyToken возвращает общее количество токенов в блокчейне
// @Summary Получение Total Supply токенов
// @Description Обращается к контракту в блокчейне и возвращает Total Supply токенов
// @Tags tokens
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 405 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /total_supply [get]
func (tc *TokenController) TotalSupplyToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	resp, err := tc.service.TotalSupplyToken(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{"result": resp}
	RespondWithJSON(w, http.StatusOK, response)
}
