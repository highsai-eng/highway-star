package operator

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/highway-star/model"
)

type TranslateOperator struct {
}

func (o TranslateOperator) Translate(srcArticles []model.Article, dstArticles *[]model.Article) error {

	sourceLanguageCode := "ko"
	targetLanguageCode := "ja"

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "highway-star"),
	})
	if err != nil {
		return err
	}

	translateClient := translate.New(sess)

	for _, el := range srcArticles {

		textInput := translate.TextInput{
			SourceLanguageCode: &sourceLanguageCode,
			TargetLanguageCode: &targetLanguageCode,
			Text:               &el.Title,
		}

		outputText, err := translateClient.Text(&textInput)
		if err != nil {
			return err
		}

		article := model.Article{Title: *outputText.TranslatedText}
		*dstArticles = append(*dstArticles, article)
	}

	return nil
}
