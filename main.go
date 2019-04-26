package main

import (
	"fmt"
	"github.com/highway-star/model"
	"github.com/highway-star/operator"
)

func main() {
	fmt.Println("Hello World!")
	scraping := operator.ScrapingOperator{}
	translator := operator.TranslateOperator{}

	rawArticles := make([]model.Article, 0)
	translatedArticles := make([]model.Article, 0)

	scraping.Operate(&rawArticles)
	translator.Operate(rawArticles, &translatedArticles)
}
