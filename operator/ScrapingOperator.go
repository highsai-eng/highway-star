package operator

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type ScrapingOperator struct {
}

func (o *ScrapingOperator) Operate() error {

	doc, err := o.fetchHtml("http://www.ilbe.com/politics")
	if err != nil {
		return err
	}

	doc.Find("tbody").Find("tr").Each(func(index int, s *goquery.Selection) {
		url, exists := s.Find(".title").Find("a").Attr("href")

		if exists {
			fmt.Printf("url : %s", url)
			fmt.Println()
		}
	})

	return nil
}

func (o *ScrapingOperator) fetchHtml(url string) (*goquery.Document, error) {

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
