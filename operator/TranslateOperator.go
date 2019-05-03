package operator

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/highway-star/model"
)

// TranslateOperator 翻訳オペレーター構造体
type TranslateOperator struct {
}

// Translate 翻訳実行
func (o TranslateOperator) Translate(srcArticle model.Article, dstArticle *model.Article) error {

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

	titleInput := translate.TextInput{
		SourceLanguageCode: &sourceLanguageCode,
		TargetLanguageCode: &targetLanguageCode,
		Text:               &srcArticle.Title,
	}

	authorInput := translate.TextInput{
		SourceLanguageCode: &sourceLanguageCode,
		TargetLanguageCode: &targetLanguageCode,
		Text:               &srcArticle.Author,
	}

	contentInput := translate.TextInput{
		SourceLanguageCode: &sourceLanguageCode,
		TargetLanguageCode: &targetLanguageCode,
		Text:               &srcArticle.Content,
	}

	titleOutput, err := translateClient.Text(&titleInput)
	if err != nil {
		return err
	}

	authorOutput, err := translateClient.Text(&authorInput)
	if err != nil {
		return err
	}

	contentOutput, err := translateClient.Text(&contentInput)
	if err != nil {
		return err
	}

	dstArticle.Title = *titleOutput.TranslatedText
	dstArticle.Author = *authorOutput.TranslatedText
	dstArticle.Content = *contentOutput.TranslatedText

	return nil
}
