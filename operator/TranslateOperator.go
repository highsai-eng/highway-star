package operator

import (
	"fmt"
	"github.com/highway-star/model"
)

type TranslateOperator struct {
}

func (o TranslateOperator) Operate(rawArticles []model.Article, translatedArticles *[]model.Article) error {

	for _, el := range rawArticles {
		fmt.Printf("title:%s", el.Title)
		fmt.Println()
	}

	return nil
}
