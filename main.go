package plugins

import (
	"math/rand"
	"strings"
)

// RandInt uses math/rand to return a random integer
func RandInt() int {
	return rand.Int()
}

// RandString generates and returns a random 50 character string
func RandString(n int) string {
	rand.Seed(int64(n))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < 50; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
