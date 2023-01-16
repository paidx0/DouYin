package models

import "time"

// UserPulishVideo 用户发布视频绑定表
type UserPulishVideo struct {
	Id        int
	UserKey   string
	VideoKey  string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table UserPulishVideo) TableName() string {
	return "userpulishvideo"
}
