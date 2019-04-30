package model

import "time"

type Article struct {
	Title          string
	Content        string
	ThumbnailImage string
	OtherImages    []string
	Tags           []string
	Published      time.Time
	Comments       []Comment
}
