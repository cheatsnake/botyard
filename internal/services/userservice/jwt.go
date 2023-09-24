package userservice

import (
	"botyard/pkg/exterr"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type userTokenClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateUserToken(userId string) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := userTokenClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		// TODO logs
		return "", time.Time{}, exterr.ErrorInternalServer("failed to create an authorization token for a new user")
	}

	return tokenString, expirationTime, nil
}

func VerifyUserToken(tokenString string) (string, error) {
	utc := &userTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, utc, func(token *jwt.Token) (any, error) {
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		// TODO logs
		return "", exterr.ErrorUnauthorized("user is unauthorized")
	}

	return utc.UserId, nil
}
