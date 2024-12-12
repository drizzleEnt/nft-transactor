package app

import (
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/drizzleent/nft-transactor/config"
	"github.com/drizzleent/nft-transactor/controller"
	"github.com/drizzleent/nft-transactor/db"
	"github.com/drizzleent/nft-transactor/repository"
	"github.com/drizzleent/nft-transactor/routes"
	"github.com/drizzleent/nft-transactor/service"
)

type App struct {
	mux *http.ServeMux

	db              *sql.DB
	tokenRepository repository.TokenRepository
	tokenService    service.TokenService
	tokenController *controller.TokenController
}

func NewApp() *App {
	a := &App{}
	return a
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(net.JoinHostPort(config.HTTP_HOST, config.HTTP_PORT), a.Mux()))
}

func (a *App) Close() {
	a.db.Close()
}

func (a *App) Mux() *http.ServeMux {
	if a.mux == nil {
		a.mux = routes.SetupRouter(a.TokenController())
	}

	return a.mux
}

func (a *App) DBClient() *sql.DB {
	if a.db == nil {
		cl, err := db.ConnectDB(config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
		if err != nil {
			log.Fatalf("Failed get data base connection: %s", err.Error())
		}
		a.db = cl
	}
	return a.db
}

func (a *App) TokenRepository() repository.TokenRepository {
	if a.tokenRepository == nil {
		a.tokenRepository = repository.NewTokenRepository(a.DBClient())
	}
	return a.tokenRepository
}

func (a *App) TokenService() service.TokenService {
	if a.tokenService == nil {
		a.tokenService = service.NewTokenService(a.TokenRepository())
	}
	return a.tokenService
}

func (a *App) TokenController() *controller.TokenController {
	if a.tokenController == nil {
		a.tokenController = controller.NewTokenController(a.TokenService())
	}
	return a.tokenController
}
