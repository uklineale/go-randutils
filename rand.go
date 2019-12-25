package randutils

import (
	"math/rand"
	"time"
)

// Charset - default charset [a-zA-Z0-9]
const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringWithCharset - generates a random string of given length
// using the given charset
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// String - generates a random string of given length
// using the default charset [a-zA-Z0-9]
func String(length int) string {
	return StringWithCharset(length, Charset)
}