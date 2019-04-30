package model

import "time"

type Comment struct {
	No        int8
	RepliedNo int8
	Name      string
	Main      bool
	Content   string
	like      int8
	dislike   int8
	Published time.Time
}
