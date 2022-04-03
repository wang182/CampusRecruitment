package utils

import (
	crand "crypto/rand"

	"math"
	"math/big"
	"math/rand"
	"time"
)

const letterAndDigit = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	n, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	if n == nil {
		n = big.NewInt(time.Now().UnixNano())
	}
	rand.Seed(n.Int64())
}

func RandomStr(n int) string {
	r := make([]byte, n)
	for i := 0; i < n; i++ {
		r[i] = letterAndDigit[rand.Intn(len(letterAndDigit))]
	}
	return string(r)
}
