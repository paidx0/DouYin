package models

import "time"

// UserFocusOn 用户关注列表
type UserFocusOn struct {
	Id        int
	UserKey   string
	ToUserKey string
	IsFollow  bool
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table UserFocusOn) TableName() string {
	return "userfocuson"
}
