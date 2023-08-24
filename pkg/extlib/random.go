package extlib

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func RandomToken(length int) (string, error) {
	tokenLen := big.NewInt(int64(len(alphabet)))
	tokenBytes := make([]byte, length)

	for i := 0; i < length; i++ {
		randIndex, err := rand.Int(rand.Reader, tokenLen)
		if err != nil {
			return "", err
		}

		tokenBytes[i] = alphabet[randIndex.Int64()]
	}

	return string(tokenBytes), nil
}

func RandomHexString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
