package operator

import (
	"fmt"
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

	// TODO: add sleep.

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch web page. url:%s", url)
	}

	// TODO: error catch.
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch web page. url:%s, code:%d, message:%s", url, res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML. url:%s", url)
	}

	return doc, nil
}