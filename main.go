package main

import (
	"log"
	"net"
	"net/http"

	"github.com/drizzleent/nft-transactor/config"
	"github.com/drizzleent/nft-transactor/db"
	"github.com/drizzleent/nft-transactor/routes"
)

func main() {
	db, err := db.ConnectDB(config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	if err != nil {
		log.Fatalf("Failed get data base connection: %s", err.Error())
	}
	defer db.Close()

	r := routes.SetupRouter(db)

	log.Fatal(http.ListenAndServe(net.JoinHostPort(config.HTTP_HOST, config.HTTP_PORT), r))
}
