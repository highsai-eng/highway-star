package main

import (
	"flag"
	"log"
	"os"

	"github.com/highway-star/model"
	"github.com/highway-star/operator"
)

func init() {
	log.SetPrefix("[highway-star]")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

var (
	keyword = flag.String("keyword", "ComfortWoman", "Specify the word to be searched in ilbe.")
)

func main() {

	log.Print("main process has started.")

	flag.Parse()

	scraper := operator.ScrapeOperator{}
	//translator := operator.TranslateOperator{}
	//uploader := operator.UploadOperator{}

	srcArticle := model.Article{}
	//srcArticles := make([]model.Article, 0)
	//dstArticles := make([]model.Article, 0)

	if err := scraper.Scrape(*keyword, &srcArticle); err != nil {
		log.Fatal(err)
	}

	log.Print(srcArticle.Uri)
	log.Print(srcArticle.Title)
	log.Print(srcArticle.Author)
	log.Print(srcArticle.Published)
	log.Print(srcArticle.Content)
	log.Print(srcArticle.ThumbnailImageUri)
	log.Print(srcArticle.ContentImageUris)
	log.Print(srcArticle.Categories)
	log.Print(srcArticle.Tags)
	log.Print(srcArticle.Comments)

	//if err := translator.Translate(srcArticles, &dstArticles); err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err := uploader.Upload(dstArticles); err != nil {
	//	log.Fatal(err)
	//}

	log.Print("main process has ended.")
}
