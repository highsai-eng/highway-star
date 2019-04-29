package model

import "time"

type Article struct {
	Title          string
	Description    string
	ThumbnailImage string
	OtherImages    []string
	Tags           []string
	Published      time.Time
	Comments       []Comment
}
