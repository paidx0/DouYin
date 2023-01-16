package utils

import (
	"DouYin/global"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaim struct {
	ID       int
	Username string
	UserKey  string
	jwt.StandardClaims
}

// CreateToken 生成token
func CreateToken(id int, userkey, name string) (string, error) {
	uc := UserClaim{
		ID:       id,
		Username: name,
		UserKey:  userkey,
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
	if _, ok := claims.Claims.(UserClaim); ok && claims.Valid {
		return uc, nil
	}
	return nil, errors.New("token失效的！！！")
}
