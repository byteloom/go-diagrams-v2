package rand

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijlkmnopqrstuvwxyz"

var seed = rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec

func String(length int) string {
	return StringWithCharset(length, charset)
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}

	return string(b)
}
