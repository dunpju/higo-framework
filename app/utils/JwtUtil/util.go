package JwtUtil

import (
	"github.com/dgrijalva/jwt-go"
	"higo-framework/app/entity/TokenPayloadEntity"
	"higo-framework/app/errcode"
	"higo-framework/app/exception/ErrorException"
	"time"
)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

const jwtSecret = "4ba85a04bbcfe9ee"

func Encrypt(tokenPayloadEntity *TokenPayloadEntity.Impl, expire time.Duration) string {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)
	claims := Claims{
		tokenPayloadEntity.UserId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))
	if nil != err {
		panic(err)
	}
	return token
}

func Decrypt(tokenString string) int64 {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if nil != err {
		panic(err)
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims.UserId
		}
	}
	ErrorException.Throw(errcode.TokenError)
	return 0
}
