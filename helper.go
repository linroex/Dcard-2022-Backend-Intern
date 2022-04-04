package main

import "math/rand"

func RandString(n int) string {
	const letterBytes = "abcdefghijkmnorstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ0123456789"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
