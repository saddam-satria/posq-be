package commons

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secret = GoDotEnv("SECRET_KEY")

func GenerateToken(id string, email string, expired int64, secretKey ...string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = expired

	if secretKey != nil {
		secret = secretKey[0]
	}

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func VerifyToken(token string) (jwt.MapClaims, error) {
	tokenString, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenString.Claims.(jwt.MapClaims); ok && tokenString.Valid {
		return claims, nil
	}

	return nil, err

}