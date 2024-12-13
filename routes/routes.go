package routes

import (
	"net/http"

	_ "github.com/drizzleent/nft-transactor/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/drizzleent/nft-transactor/controller"
)

func SetupRouter(tc *controller.TokenController) *http.ServeMux {
	r := http.NewServeMux()

	// TOKENS
	r.HandleFunc("/tokens/create", func(w http.ResponseWriter, r *http.Request) {
		tc.CreateToken(w, r)
	})

	r.HandleFunc("/tokens/list", func(w http.ResponseWriter, r *http.Request) {
		tc.ListToken(w, r)
	})

	r.HandleFunc("/tokens/total_supply", func(w http.ResponseWriter, r *http.Request) {
		tc.TotalSupplyToken(w, r)
	})
	//END TOKENS

	r.Handle("/swagger/*", httpSwagger.WrapHandler)
	return r
}
