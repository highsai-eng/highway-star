package main

import (
	"os"

	"github.com/highway-star/model"
	"github.com/highway-star/operator"
)

func main() {

	scraping := operator.ScrapingOperator{}
	translator := operator.TranslateOperator{}
	uploader := operator.UploadOperator{}

	srcArticles := make([]model.Article, 0)
	dstArticles := make([]model.Article, 0)

	if err := scraping.Scraping(&srcArticles); err != nil {
		// TODO: log export
		os.Exit(1)
	}

	if err := translator.Translate(srcArticles, &dstArticles); err != nil {
		// TODO: log export
		os.Exit(2)
	}

	if err := uploader.Upload(); err != nil {
		// TODO: log export
		os.Exit(3)
	}
}
