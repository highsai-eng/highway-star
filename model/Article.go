package model

import "time"

type Article struct {
	Uri            string
	Title          string
	Author         string
	Published      time.Time
	Content        string
	ThumbnailImage string
	OtherImages    []string
	Tags           []string
	Comments       []Comment
}
