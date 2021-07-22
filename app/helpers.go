package app

import (
	"encoding/json"
	"log"
	"net/http"
	"goarticle/app/models"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}


func mapArticleToJSON(p *models.Article) models.JsonArticle {
	return models.JsonArticle{
		Id:      p.Id,
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
	}
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}