package models

type Article struct {
	Id		int64	`db:"id"`
	Title	string	`db:"title"`
	Content	string	`db:"content"`
	Author 	string	`db:"author"`
}

// for parsing json
type JsonArticle struct {
	Id		int64	`json:"id"`
	Title	string	`json:"title"`
	Content	string	`json:"content"`
	Author 	string	`json:"author"`
}

// for request
type JsonRequest struct {
	Title	string	`json:"title"`
	Content	string	`json:"content"`
	Author  string `json:"author"`
}