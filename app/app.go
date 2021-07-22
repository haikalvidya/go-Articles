package app

import (
	"github.com/gorilla/mux"
	"goarticle/app/database"
)

type App struct	{
	Router *mux.Router
	DB		database.ArticleDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()). Methods("GET")
	a.Router.HandleFunc("/api/articles", a.CreateArticleHandler()).Methods("POST")
	a.Router.HandleFunc("/api/articles", a.GetArticlesHandler()).Methods("GET")
}

