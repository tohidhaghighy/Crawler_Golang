package services

import (
	"crawler/models"
)

func InsertTOMongodb(article models.Article) (status bool) {
	return true
}

func FindDataMongo(article string) (articles []models.Article) {
	return nil
}
