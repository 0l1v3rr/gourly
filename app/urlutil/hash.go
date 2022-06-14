package urlutil

import "math/rand"

const chars string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomUrl(length int) string {
	res := ""

	for i := 0; i <= length; i++ {
		res += string(chars[rand.Intn(len(chars)-1)])
	}

	return res
}
