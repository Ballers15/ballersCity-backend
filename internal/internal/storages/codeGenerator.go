package storages

import (
	"crypto/rand"
	"math/big"
)

const DefaultPasswordLength = 8

func GenerateOneTimePassword(n int) (string, error) {
	if n < 0 {
		n = DefaultPasswordLength
	}
	const symbols = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		if err != nil {
			return "", err
		}
		result[i] = symbols[num.Int64()]
	}
	return string(result), nil
}
