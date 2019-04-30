package operator

import (
	"encoding/base64"
	"fmt"

	"github.com/highway-star/model"
	"gopkg.in/resty.v1"
)

type UploadOperator struct {
}

const (
	username = "kaito.higa"
	password = "ZhHt SP54 40bH Wddw KsvY 2xWP"
)

func (o *UploadOperator) Upload(articles []model.Article) error {

	authInfo := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	resp, err := resty.R().SetHeaders(map[string]string{
		"Authorization": "Basic " + authInfo,
		"Content-Type":  "application/json",
	}).SetBody(map[string]interface{}{
		"title":   "ポストテスト",
		"status":  "draft",
		"content": "テスト",
		//"categories": 1,
		//"tags":       "1,2,3",
		//"date":       "2019-04-30T10:00:00",
		//"excerpt":    "Read this awesome post",
		//"password":   "12$45",
		//"slug":       "new-test-post",
	}).Post("http://nida.xsrv.jp/wp-json/wp/v2/posts")

	// explore response object
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())

	return nil
}
