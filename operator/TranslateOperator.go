package operator

import (
	"fmt"

	"github.com/highway-star/model"
)

type TranslateOperator struct {
}

func (o TranslateOperator) Translate(srcArticles []model.Article, dstArticles *[]model.Article) error {

	for _, el := range srcArticles {
		fmt.Printf("title:%s", el.Title)
		fmt.Println()
	}

	return nil
}
