package app

import (
	"fmt"
	"net/http"
	"goarticle/app/models"
	"log"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Post API")
	}
}

func (a *App) CreateArticleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.JsonRequest{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		// Create the post
		p := &models.Article{
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		// Save in DB
		err = a.DB.CreateArticle(p)
		if err != nil {
			log.Printf("Cannot save post in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapArticleToJSON(p)
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) GetArticlesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := a.DB.GetArticle()
		if err != nil {
			log.Printf("Cannot get posts, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonArticle, len(articles))
		for idx, article := range articles {
			resp[idx] = mapArticleToJSON(article)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}