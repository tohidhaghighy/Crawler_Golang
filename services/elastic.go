package services

import (
	"crawler/models"
	"time"
)

// if search type==0 search with title
// if search type==1 search with body
// if search type==2 search with date
// if search type==3 search in all title and body
type SearchModel struct {
	title       string
	body        string
	date        time.Time
	search_type int
}

func InsertToElastic(article models.Article) (status bool) {
	return true
}

func GetFromElastic(article SearchModel) (articles []models.Article) {
	return nil
}
