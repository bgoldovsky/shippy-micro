package main

import (
	pb "github.com/bgoldovsky/shippy-micro/userserver/proto/user"
	"github.com/dgrijalva/jwt-go"
)

var key = []byte("MyNewSecretKey99999")

// UserClaim ...
type UserClaim struct {
	User *pb.User
	jwt.StandardClaims
}

// Authable ...
type Authable interface {
	Encode(user *pb.User) (string, error)
	Decode(token string) (*UserClaim, error)
}

// TokenService ...
type TokenService struct {
	repo Repository
}

// Encode ...
func (ts *TokenService) Encode(user *pb.User) (string, error) {
	claim := UserClaim{
		user,
		jwt.StandardClaims{
			ExpiresAt: 20000,
			Issuer:    "user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(key)
}

// Decode ...
func (ts *TokenService) Decode(token string) (*UserClaim, error) {
	tokenType, err := jwt.ParseWithClaims(string(key), &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if claim, isOk := tokenType.Claims.(*UserClaim); isOk && tokenType.Valid {
		return claim, nil
	}
	return nil, err
}
