package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte = []byte(os.Getenv("SECRET_KEY"))

type DataClaims struct {
	ID    uint 		`json:"id"`
	Email string 	`json:"email"`
	Admin bool 		`json:"admin"`
}

type Claims struct {
	DataClaims
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, email string, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		DataClaims: DataClaims{
			ID: userID,
			Email: email,
			Admin: isAdmin,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "Failed to sign token", err
	}

	return signedToken, nil
}

func ValidateToken (requestToken string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(requestToken, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
