package main

import (
	"flag"
	"log"
	"os"

	"github.com/highway-star/model"
	"github.com/highway-star/operator"
)

var (
	keyword = flag.String("keyword", "ComfortWoman", "Specify the word to be searched in ilbe.")
)

func init() {
	log.SetPrefix("[highway-star]")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func main() {

	log.Print("main process has started.")

	flag.Parse()

	scraper := operator.ScrapeOperator{}
	translator := operator.TranslateOperator{}
	//uploader := operator.UploadOperator{}

	srcArticle := model.Article{}
	dstArticle := model.Article{}

	if err := scraper.Scrape(*keyword, &srcArticle); err != nil {
		log.Fatal(err)
	}

	log.Print(srcArticle.URI)
	log.Print(srcArticle.Title)
	log.Print(srcArticle.Author)
	log.Print(srcArticle.Published)
	log.Print(srcArticle.Content)
	log.Print(srcArticle.ThumbnailImageURI)
	log.Print(srcArticle.ContentImageURIs)
	log.Print(srcArticle.Categories)
	log.Print(srcArticle.Tags)
	log.Print(srcArticle.Comments)

	if err := translator.Translate(srcArticle, &dstArticle); err != nil {
		log.Fatal(err)
	}

	log.Print(dstArticle.URI)
	log.Print(dstArticle.Title)
	log.Print(dstArticle.Author)
	log.Print(dstArticle.Published)
	log.Print(dstArticle.Content)
	log.Print(dstArticle.ThumbnailImageURI)
	log.Print(dstArticle.ContentImageURIs)
	log.Print(dstArticle.Categories)
	log.Print(dstArticle.Tags)
	log.Print(dstArticle.Comments)

	//if err := uploader.Upload(dstArticles); err != nil {
	//	log.Fatal(err)
	//}

	log.Print("main process has ended.")
}
