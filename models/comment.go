package models

import "time"

// Comment 评论信息表
type Comment struct {
	Id          int
	UserKey     string
	VideoKey    string
	CommentText string
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
	DeletedAt   time.Time `xorm:"deleted"`
}

func (table Comment) TableName() string {
	return "comment"
}
