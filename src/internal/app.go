package internal

import (
	"menu/internal/fts"
	"menu/internal/repository"
	"menu/internal/transport"
	"menu/pkg/manticore"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type App struct {
	handelrs struct {
		search *transport.MenuSearchHandler
	}
}

func InitNewApp(db *sqlx.DB, manticoreClient *manticore.ManticoreClient) *App {
	menuRepo := repository.NewMenuRespository(db)
	menuSearch := fts.NewMenuSearch(manticoreClient)

	app := &App{}
	app.handelrs.search = transport.NewMenuSearchHandler(menuRepo, menuSearch)

	return app
}

func (a *App) GetRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/search", a.handelrs.search.GetRouter())

	return router
}
