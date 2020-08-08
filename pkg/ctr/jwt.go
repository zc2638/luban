/**
 * Created by zc on 2020/6/9.
 */
package ctr

import (
	"github.com/dgrijalva/jwt-go"
	"luban/pkg/errs"
)

type JwtClaims struct {
	jwt.StandardClaims
	User JwtUserInfo
}

type JwtUserInfo struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

// JwtCreate returns the JWT token by claims and secret
func JwtCreate(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// JwtParse returns the claims by JWT token and secret
func JwtParse(claims jwt.Claims, tokenStr string, secret string) error {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.New("Token: invalid signature")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errs.New("Invalid Token")
	}
	return nil
}
