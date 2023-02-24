package util

import (
	"math/rand"
	"strings"
	"time"
)

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func RandomString(length int) string {
	var sb strings.Builder
	k := len(letters)

	for i := 0; i < length; i++ {
		ch := letters[rand.Intn(k)]
		sb.WriteByte(ch)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(1, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "RUB"}
	return currencies[rand.Intn(len(currencies))]
}
