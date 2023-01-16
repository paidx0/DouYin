package models

import "time"

// UserLikeVideo 用户点赞视频绑定表
type UserLikeVideo struct {
	Id        int
	UserKey   string
	VideoKey  string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table UserLikeVideo) TableName() string {
	return "userlikevideo"
}
