package models

import "time"

// Video 视频信息表
type Video struct {
	Id            int
	Title         string
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int
	CommentCount  int
	Name          string
	Tag           string
	VideoKey      string
	UserKey       string
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
	DeletedAt     time.Time `xorm:"deleted"`
}

func (table Video) TableName() string {
	return "video"
}
