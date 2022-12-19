package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"log"
)

var sign = "CuCreateTR"

func NewToken(claims jwt.MapClaims) (string, error) {
	signedString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(sign))
	if err != nil {
		log.Printf("new token error:%v", err)
		return "", err
	}
	return signedString, nil
}

func ParseToken(token string) (map[string]interface{}, error) {
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// 使用签名解析用户传入的token,获取载荷部分数据
		return []byte(sign), nil
	})
	if err != nil {
		log.Printf("new token error:%v", err)
		return make(map[string]interface{}), err
	}
	claims := tk.Claims.(jwt.MapClaims)
	return claims, nil
}
