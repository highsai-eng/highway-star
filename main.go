package main

import (
	"fmt"
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

func main() {

	log.Print("main process has started.")

	scraper := operator.ScrapeOperator{}
	translator := operator.TranslateOperator{}
	uploader := operator.UploadOperator{}

	srcArticles := make([]model.Article, 0)
	dstArticles := make([]model.Article, 0)

	if err := scraper.Scrape(&srcArticles); err != nil {
		log.Fatal(err)
	}

	if err := translator.Translate(srcArticles, &dstArticles); err != nil {
		log.Fatal(err)
	}

	if err := uploader.Upload(dstArticles); err != nil {
		log.Fatal(err)
	}

	for _, el := range srcArticles {
		fmt.Printf("befor:%s", el.Title)
		fmt.Println()
	}

	for _, el := range dstArticles {
		fmt.Printf("after:%s", el.Title)
		fmt.Println()
	}

	log.Print("main process has ended.")
}
