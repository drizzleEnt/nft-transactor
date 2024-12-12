package models

type Token struct {
	UniqueHash string
	TxHash     string
	MediaUrl   string
	Owner      string
}

type CreateTokenRequest struct {
	MediaUrl string `json:"media_url"`
	Owner    string `json:"owner"`
}
