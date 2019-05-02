package model

import "time"

// Article 記事構造体
type Article struct {
	URI               string
	Title             string
	Author            string
	Published         time.Time
	Content           string
	ThumbnailImageURI string
	ContentImageURIs  []string
	Categories        []string
	Tags              []string
	Comments          []Comment
}
