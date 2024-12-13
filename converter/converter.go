package converter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/drizzleent/nft-transactor/models"
)

func FromRequestToCreateTokenRequest(r *http.Request) (*models.CreateTokenRequest, error) {
	var req models.CreateTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshal request: %s", err.Error())
	}

	return &req, nil
}

func FromRequestParamsToModel(r *http.Request) *models.RequestParams {

	var params models.RequestParams

	limitString := r.URL.Query().Get("limit")
	if limitString == "" {
		params.Limit = 200
	} else {
		lim, err := strconv.Atoi(limitString)
		if err != nil {
			lim = 200
		} else {
			if lim < 200 {
				lim = 200
			}
			if lim > 500 {
				lim = 500
			}
			params.Limit = lim
		}

	}
	offsetString := r.URL.Query().Get("offset")
	if offsetString == "" {
		params.Offset = 0
	} else {
		offset, err := strconv.Atoi(offsetString)
		if err != nil {
			offset = 0
		}
		params.Offset = offset
	}
	return &params
}
