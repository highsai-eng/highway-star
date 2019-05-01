package operator

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/highway-star/constant"

	"github.com/PuerkitoBio/goquery"
	"github.com/highway-star/model"
)

type ScrapeOperator struct {
}

const (
	maxReadPage   = 5
	minReplyCount = 10
	layout        = "2006.01.02 15:04:05"
)

func (o *ScrapeOperator) Scrape(keyword string, article *model.Article) error {

	found := false

	for i := 0; i < maxReadPage; i++ {

		if found {
			break
		}

		doc, err := o.fetchHtml(o.generateSearchUrl(constant.Get().Keywords[keyword].Korean, i))
		if err != nil {
			return err
		}

		doc.Find(".searchResult").Find("li").Each(func(index int, s *goquery.Selection) {

			if found {
				return
			}

			_, exists := s.Find("img").Attr("src")

			if exists {
				reply, _ := strconv.Atoi(
					s.Find("dl").Find("dt").Find("span").Find("em").Text())

				if reply >= minReplyCount {
					url, _ := s.Find("dl").Find("dt").Find("a").Attr("href")
					if err := o.analyzeArticle(url, article); err != nil {
						log.Print(err)
					}

					found = true
				}
			}
		})
	}

	return nil
}

func (o *ScrapeOperator) analyzeArticle(url string, article *model.Article) error {

	doc, err := o.fetchHtml(url)
	if err != nil {
		return err
	}

	var (
		content          string
		contentImageUris []string
	)

	readHeader := doc.Find("div.originalContent").Find("div.readHeader")
	readBody := doc.Find("div.originalContent").Find("div.readBody")

	uri := strings.TrimSpace(readHeader.Find("div.uri").Find("a").Text())
	title := strings.TrimSpace(readHeader.Find("div.title").Text())
	author := strings.TrimSpace(readHeader.Find("div.userInfo").Find("span").Text())
	dateStr := strings.TrimSpace(readHeader.Find("div.date").Text())

	readBody.Find("div#copy_layer_1").Find("p").Each(func(index int, s *goquery.Selection) {

		_, exists := s.Find("img").Attr("src")

		if exists {
			s.Find("img").Each(func(index int, s *goquery.Selection) {
				imageUri, _ := s.Attr("src")
				contentImageUris = append(contentImageUris, imageUri)
			})
		} else {
			if strings.TrimSpace(s.Text()) != "" {
				content += strings.TrimSpace(s.Text()) + "¥n"
			}
		}
	})

	published, err := time.Parse(layout, dateStr)

	if err != nil {
		return err
	}

	article.Uri = uri
	article.Title = title
	article.Author = author
	article.Published = published
	article.Content = content
	article.ThumbnailImageUri = contentImageUris[0]
	article.ContentImageUris = contentImageUris

	// TODO:category and tag implements.
	article.Categories = []string{}
	article.Tags = []string{}

	// TODO:comment implements.
	article.Comments = []model.Comment{}

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

func (o *ScrapeOperator) generateSearchUrl(keyword string, index int) string {

	return "http://www.ilbe.com/index.php?" +
		"act=IS" +
		"&where=document" +
		"&is_keyword=" + keyword +
		"&mid=index" +
		"&search_target=title" +
		"&page=" + strconv.Itoa(index+1)
}
