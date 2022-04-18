package jwt

import (
	"time"

	jt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jt.StandardClaims
}

var secret = []byte("zhaoyunxing")

func GenerateToken(username string) (string, error) {
	user := &Claims{Username: username}
	user.ExpiresAt = time.Now().Add(time.Minute * 30).Unix()
	user.Issuer = "zhaoyunxing"
	claims := jt.NewWithClaims(jt.SigningMethodHS256, user)
	return claims.SignedString(secret)
}

func ParseToken(token string) (*Claims, error) {
	user := &Claims{}

	claims, err := jt.ParseWithClaims(token, user, func(token *jt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims != nil && claims.Valid {
		return user, err
	}
	return nil, err
}
