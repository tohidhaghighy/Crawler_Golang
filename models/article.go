package models

import "time"

type Article struct {
	Id      int
	Title   string
	Content string
	Date    time.Time
	Website string
}

func (a Article) GetTitle() string {
	return a.Title
}

func (a Article) GetContent() string {
	return a.Content
}

func (a Article) CreateArticle() Article {
	return a
}
