package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuiopasdfghjklzxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandonInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandonString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]

		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
		return RandonString(6)
}

func RandomMoney() int64 {
	return RandonInt(1, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}