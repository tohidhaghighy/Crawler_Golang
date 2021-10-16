package main

import (
	mdl "crawler/models"
	"fmt"
	"time"
)

func main() {
	article_one := mdl.Article{
		Title:   "tohid",
		Content: "haghighi",
		Date:    time.Now(),
		Id:      1,
		Website: "www.tohid.com",
	}
	fmt.Printf("%s article content \n", article_one.GetTitle())
}
