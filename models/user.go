package models

import "time"

// User 用户信息表
type User struct {
	Id            int
	Username      string
	Password      string
	FollowCount   int
	FollowerCount int
	LikeTotal     int
	UserKey       string
	CreatedAt     time.Time `xorm:"created"`
	UpdatedAt     time.Time `xorm:"updated"`
	DeletedAt     time.Time `xorm:"deleted"`
}

func (table User) TableName() string {
	return "user"
}
