package operator

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/highway-star/model"
)

type ScrapingOperator struct {
}

func (o *ScrapingOperator) Scraping(articles *[]model.Article) error {

	listDoc, err := o.fetchHtml("http://www.ilbe.com/politics")
	if err != nil {
		return err
	}

	listDoc.Find("tbody").Find("tr").Each(func(index int, s *goquery.Selection) {

		url, exists := s.Find(".title").Find("a").Attr("href")

		if exists && strings.HasPrefix(url, "http") {

			detailDoc, _ := o.fetchHtml(url)

			article := model.Article{Title: strings.TrimSpace(detailDoc.Find("div.title").Text())}
			*articles = append(*articles, article)
		}
	})

	return nil
}

func (o *ScrapingOperator) fetchHtml(url string) (*goquery.Document, error) {

	log.Printf("start fetching HTML. url:%s", url)

	// TODO: add sleep.
	//time.Sleep(60 * time.Second)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch web page. url:%s", url)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			// TODO: log export
		}
	}()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch web page. url:%s, code:%d, message:%s", url, res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML. url:%s", url)
	}

	log.Printf("end fetching HTML.")

	return doc, nil
}
