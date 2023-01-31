package utils

import (
	"errors"
	"regexp"
)

func CheckUserLayout(username, passwd string) (err error) {
	if username == "" || passwd == "" {
		return errors.New("用户名和密码不能为空")
	}

	if len(passwd) < 6 {
		return errors.New("密码长度最少6个字符")
	}

	if ok, _ := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", username); !ok {
		return errors.New("用户名不符合Email格式，或出现非法字符")
	}

	if len(username) > 32 || len(passwd) > 32 {
		return errors.New("用户名或密码最长32个字符")
	}

	return nil
}
