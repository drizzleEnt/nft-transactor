package main

import (
	"github.com/drizzleent/nft-transactor/app"
)

// @title NFT API
// @version 0.1
// description API для работы с nft токенами
// @host localhost:8080
// @BasePath /tokens
func main() {
	a := app.NewApp()
	defer a.Close()
	a.Run()
}
