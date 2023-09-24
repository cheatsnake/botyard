package userservice

import (
	"botyard/internal/entities/user"
	"testing"
)

func TestUserToken(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		newUser, err := user.New("test")
		if err != nil {
			t.Errorf("got: %v,\nexpected: %v\n", err.Error(), nil)
		}

		token, _, err := GenerateUserToken(newUser.Id, newUser.Id)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", token, err.Error(), nil)
		}

		userClaims, err := VerifyUserToken(token)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", token, err.Error(), nil)
		}

		if userClaims.Id != newUser.Id {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", token, userClaims.Id, newUser.Id)
		}
	})

	t.Run("invalid token", func(t *testing.T) {
		token := "jwt malformed"

		_, err := VerifyUserToken(token)
		if err == nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", token, err, "error")
		}
	})
}
