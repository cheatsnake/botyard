package userservice

import (
	"botyard/internal/tools/ulid"
	"testing"
)

func TestUserToken(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		userId := ulid.New()
		token, _, err := GenerateUserToken(userId)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", token, err.Error(), nil)
		}

		gotUserId, err := VerifyUserToken(token)
		if err != nil {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", token, err.Error(), nil)
		}

		if gotUserId != userId {
			t.Errorf("%#v\ngot: %v,\nexpected: %v\n", token, gotUserId, userId)
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
