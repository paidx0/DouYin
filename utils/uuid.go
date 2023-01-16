package utils

import (
	uuid "github.com/satori/go.uuid"
)

// UUID 生成唯一编号
func UUID() string {
	return uuid.NewV4().String()
}
