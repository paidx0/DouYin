package utils

import (
	"DouYin/global"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaim struct {
	Id       int
	Username string
	Userkey  string
	jwt.StandardClaims
}

// CreateToken 生成token
func CreateToken(id int, userkey, name string) (string, error) {
	uc := UserClaim{
		Id:       id,
		Username: name,
		Userkey:  userkey,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(global.CONFIG.JWT.ExpiresTime)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(global.CONFIG.JWT.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// DecodeToken 解析token
func DecodeToken(token string) (*UserClaim, error) {
	uc := new(UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.CONFIG.JWT.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := claims.Claims.(*UserClaim); ok && claims.Valid {
		return uc, nil
	}
	return nil, errors.New("token失效")
}

// CheckToken 检查UserToken
func CheckToken(token string) (uc *UserClaim, err error) {
	if token == "" {
		return nil, errors.New("token空值")
	}
	uc, err = DecodeToken(token)
	if err != nil {
		return nil, errors.New("token错误")
	}
	return
}
