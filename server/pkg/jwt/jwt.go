package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(userID string, jwtSecret string) (string, error) {

	claims := jwt.MapClaims{
		"sub": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
