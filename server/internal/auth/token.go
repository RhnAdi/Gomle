package auth

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type JWTClaim struct {
	ID    string
	Email string
	jwt.StandardClaims
}

func GenerateJWT(id string, email string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error: can't read env file, errorMessage: %F", err)
	}
	jwtKey := []byte(os.Getenv("JWT_KEY"))
	expiredTime := time.Now().Add(2 * time.Hour)
	claims := &JWTClaim{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func ValidateToken(tokenString string) (JWTClaim, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error: can't read env file, errorMessage: %F", err)
	}
	jwtKey := []byte(os.Getenv("JWT_KEY"))
	var data JWTClaim

	token, err := jwt.ParseWithClaims(
		tokenString,
		&data,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return JWTClaim{}, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return JWTClaim{}, errors.New("could't parse claims")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return JWTClaim{}, errors.New("token expired")
	}
	return data, nil
}
