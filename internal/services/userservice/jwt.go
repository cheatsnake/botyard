package userservice

import (
	"os"
	"time"

	"github.com/cheatsnake/botyard/internal/config"
	"github.com/cheatsnake/botyard/internal/entities/user"
	"github.com/cheatsnake/botyard/pkg/exterr"

	"github.com/golang-jwt/jwt/v5"
)

type userTokenClaims struct {
	user.User
	jwt.RegisteredClaims
}

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateUserToken(userId, nick string) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Duration(config.Global.Limits.User.AuthTokenLifetime) * time.Minute)
	claims := userTokenClaims{
		user.User{
			Id:       userId,
			Nickname: nick,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		// TODO logs
		return "", time.Time{}, exterr.ErrorInternalServer("Failed to create an authorization token for a new user.")
	}

	return tokenString, expirationTime, nil
}

func VerifyUserToken(tokenString string) (*userTokenClaims, error) {
	utc := &userTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, utc, func(token *jwt.Token) (any, error) {
		return jwtSecretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, exterr.ErrorUnauthorized("User is unauthorized.")
	}

	return utc, nil
}
