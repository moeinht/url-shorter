package hash

import "math/rand"

func GenerateCode() string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	code := make([]byte, 6)
	for i := 0; i < 6; i++ {
		code[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(code)
}
