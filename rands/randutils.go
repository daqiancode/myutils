package rands

import (
	"math/rand"
)

// func init() {
// 	rand.Seed(rand.Int63())
// }

var (
	AlphabetNumber       = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	AlphabetNumberNoCase = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	Numbers              = []byte("0123456789")
	Hex                  = []byte("abcdef0123456789")
)

func RandomNumber(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = Numbers[rand.Intn(len(Numbers))]
	}
	return string(b)
}

func Random(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = AlphabetNumberNoCase[rand.Intn(len(AlphabetNumberNoCase))]
	}
	return string(b)
}

func RandomLower(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = AlphabetNumberNoCase[rand.Intn(len(AlphabetNumberNoCase))]
	}
	return string(b)
}

func RandomHex(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = Hex[rand.Intn(len(Hex))]
	}
	return string(b)
}
