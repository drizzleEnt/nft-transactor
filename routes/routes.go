package routes

import (
	"database/sql"
	"net/http"

	"github.com/drizzleent/nft-transactor/controller"
)

func SetupRouter(db *sql.DB) *http.ServeMux {
	r := http.NewServeMux()

	// TOKENS
	r.HandleFunc("/tokens/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {

			return
		}
		tokenController := controller.NewTokenController()
		tokenController.CreateToken(w, r)
	})

	r.HandleFunc("/tokens/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {

			return
		}
		tokenController := controller.NewTokenController()
		tokenController.CreateToken(w, r)
	})

	r.HandleFunc("/tokens/total_supply", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {

			return
		}
		tokenController := controller.NewTokenController()
		tokenController.CreateToken(w, r)
	})
	//END TOKENS

	return r
}
