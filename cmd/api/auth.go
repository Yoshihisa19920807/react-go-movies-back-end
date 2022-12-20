package main

import (
	"fmt"
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

type JwtUser struct {
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

func (j *Auth) GenerateTokenPair(user *JwtUser) (TokenPairs, error) {
	// Create a token
	token := jwt.New(jwt.SigningMethodES256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	// these key names with 3 letters seem to be official abbreviations
	claims["sub"] = fmt.Sprint(user.ID) // subject
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix() // issued at
	claims["typ"] = "JWT"

	// Set the expiry for JWT
	claims["exp"] = time.Now().Add(j.TokenExpiryIn).Unix()

	// Create a signed token
	signedAccessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	// Create refresh token and set claims
	refreshToken := jwt.New(jwt.SigningMethodES256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["sub"] = fmt.Sprint(user.ID)
	refreshClaims["iat"] = time.Now().UTC().Unix()

	// Set the expiry for the refresh token
	refreshClaims["exp"] = time.Now().Add(j.RefreshExpiray).Unix()

	// Create signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}

	// Create tokenPairs and populate with signed tokens
	var tokenPairs = TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	// Return TokenPairs
	return tokenPairs, nil
}