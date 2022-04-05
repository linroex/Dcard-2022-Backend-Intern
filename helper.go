package main

import (
	"math/rand"
	"time"
)

var getRand = rand.Intn

func RandString(n int) string {
	const letterBytes = "abcdefghijkmnorstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixMicro())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[getRand(len(letterBytes))]
	}
	return string(b)
}
