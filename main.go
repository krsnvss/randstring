package randstring

import (
	"math/rand"
	"time"
)

// Charset for random string
const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789,.:;()<>?!@#$%&"

// Copy of Rand
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// MakeRandomString - make random string of certain length
func MakeRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
