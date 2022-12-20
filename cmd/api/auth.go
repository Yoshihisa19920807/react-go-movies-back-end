package main

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	Issuer         string
	Audience       string
	Secret         string
	TokenExpiryIn  time.Duration
	RefreshExpiray time.Duration
	CookieDomain   string
	CookiePath     string
	CookieName     string
}

type JwtUset struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}
