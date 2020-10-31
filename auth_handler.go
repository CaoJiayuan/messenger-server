package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/enorith/framework/http"
	"github.com/enorith/framework/http/content"
	"github.com/enorith/framework/http/contracts"
	"github.com/enorith/supports/str"
	"time"
)

type AuthMiddleware struct {
	JwtKey string
}

func (a AuthMiddleware) Handle(r contracts.RequestContract, next http.PipeHandler) contracts.ResponseContract {
	token, err := r.BearerToken()
	if err != nil {
		return content.ErrResponseFromError(err, 401, nil)
	}
	claims := jwt.StandardClaims{}
	t, e := jwt.ParseWithClaims(string(token), &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.JwtKey), nil
	})
	if e != nil {
		return content.ErrResponseFromError(e, 401, nil)
	}

	if !t.Valid {
		return content.ErrResponseFromError(errors.New("invalid token"), 401, nil)
	}

	validErr := claims.Valid()
	if validErr != nil {
		return content.ErrResponseFromError(validErr, 401, nil)
	}


	return next(r)
}

type Token struct {
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}

type Credentials struct {
	content.Request
	Key string `input:"key"`
	Id  string `input:"_id"`
}

type AuthHandler struct {
	MasterKey       string
	JwtExpireSecond int
	JwtKey          string
}

func (a AuthHandler) Login(c Credentials) (contracts.ResponseContract, error) {
	now := time.Now().Unix()

	exp := now + int64(a.JwtExpireSecond)

	if c.Key == a.MasterKey {
		claims := jwt.MapClaims{
			"iss": "",
		}
		claims["jti"] = c.Id
		claims["sub"] = c.Id
		claims["iat"] = now
		claims["exp"] = exp
		claims["aud"] = str.RandString(16)
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

		tokenString, err := token.SignedString([]byte(a.JwtKey))
		if err != nil {
			return nil, err
		}

		return content.JsonResponse(Token{
			Token: tokenString,
			Exp:   exp,
		}, 200, nil), nil
	}

	return content.JsonResponse(Status(401), 401, nil), nil
}
