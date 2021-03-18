package utils

import (
	"github.com/dgrijalva/jwt-go"
	"shop/config"
	"time"
)

var jwtSecret = []byte(config.JwtSecret)

type Claims struct {
	UID uint `json:"uid"`
	jwt.StandardClaims
}

//封装Token
func GenerateToken(uid uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		uid,
		jwt.StandardClaims{
			//过期时间设置为3个小时后
			ExpiresAt: expireTime.Unix(),
			Issuer:    "谢金武",
		},
	}
	//加密Token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//签名
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//解析Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	//fmt.Println(err)
	return nil, err
}
