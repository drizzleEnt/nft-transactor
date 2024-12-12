package converter

import (
	"encoding/json"
	"fmt"
	"net/http"

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
