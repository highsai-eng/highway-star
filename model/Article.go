package model

import "time"

type Article struct {
	Uri               string
	Title             string
	Author            string
	Published         time.Time
	Content           string
	ThumbnailImageUri string
	ContentImageUris  []string
	Categories        []string
	Tags              []string
	Comments          []Comment
}
