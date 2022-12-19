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
		return []byte(sign), nil
	})
	if err != nil {
		log.Printf("parse token error: %v", err)
		return nil, err
	}
	claims := tk.Claims.(jwt.MapClaims)
	return claims, nil
}
