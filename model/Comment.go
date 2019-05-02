package model

import "time"

// Comment コメント構造体
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
