package operator

import (
	"encoding/base64"
	"log"

	"github.com/highway-star/model"
	"gopkg.in/resty.v1"
)

// UploadOperator アップロードオペレーター構造体
type UploadOperator struct {
}

const (
	username          = "kaito.higa"
	password          = "8xyc WFyN o9Pi Cgy5 oujp uX3U"
	wordPressEndpoint = "http://nida.xsrv.jp/wp-json/wp/v2/posts"
)

// Upload アップロード実行
func (o *UploadOperator) Upload(articles []model.Article) error {

	authInfo := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	for _, element := range articles {
		if err := o.post(authInfo, element); err != nil {
			return err
		}
	}

	return nil
}

func (o *UploadOperator) post(authInfo string, article model.Article) error {

	header := map[string]string{
		"Authorization": "Basic " + authInfo,
		"Content-Type":  "application/json",
	}

	body := map[string]interface{}{
		"title":   article.Title,
		"status":  "publish",
		"content": "テスト",
		//"categories": 1,
		//"tags":       "1,2,3",
		//"date":       "2019-04-30T10:00:00",
		//"excerpt":    "Read this awesome post",
		//"password":   "12$45",
		//"slug":       "new-test-post",
	}

	resp, err := resty.R().SetHeaders(header).SetBody(body).Post(wordPressEndpoint)

	if err != nil {
		return err
	}

	log.Printf("\nError: %v", err)
	log.Printf("\nResponse Status Code: %v", resp.StatusCode())
	log.Printf("\nResponse Status: %v", resp.Status())
	log.Printf("\nResponse Time: %v", resp.Time())
	log.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	log.Printf("\nResponse Body: %v", resp)

	return nil
}
