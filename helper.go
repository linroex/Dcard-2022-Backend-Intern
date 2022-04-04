package main

import (
	"math/rand"
	"time"
)

func RandString(n int) string {
	const letterBytes = "abcdefghijkmnorstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixMicro())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
