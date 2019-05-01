package operator

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/highway-star/model"
)

type ScrapeOperator struct {
}

const (
	maxReadPage   = 5
	minReplyCount = 10
)

func (o *ScrapeOperator) Scrape(keyword string, article *model.Article) error {

	found := false

	for i := 0; i < maxReadPage; i++ {

		if found {
			break
		}

		listDoc, err := o.fetchHtml(
			"http://www.ilbe.com/index.php?" +
				"act=IS" +
				"&where=document" +
				"&is_keyword=" + keyword +
				"&mid=index" +
				"&search_target=title" +
				"&page=" + strconv.Itoa(i+1))

		if err != nil {
			return err
		}

		listDoc.Find(".searchResult").Find("li").Each(func(index int, s *goquery.Selection) {

			if found {
				return
			}

			_, exists := s.Find("img").Attr("src")

			if exists {
				reply, _ := strconv.Atoi(
					s.Find("dl").Find("dt").Find("span").Find("em").Text())

				if reply >= minReplyCount {
					url, _ := s.Find("dl").Find("dt").Find("a").Attr("href")
					o.analyzeArticle(url, article)

					found = true
				}
			}
		})
	}

	return nil
}

func (o *ScrapeOperator) analyzeArticle(url string, article *model.Article) error {

	log.Print(url)

	return nil
}

func (o *ScrapeOperator) fetchHtml(url string) (*goquery.Document, error) {

	log.Printf("start fetching HTML. url:%s", url)

	// TODO: add sleep.
	//time.Sleep(60 * time.Second)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch web page. url:%s", url)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Print(err)
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
