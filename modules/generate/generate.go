package generate

import "math/rand"

func GenerateString(n int) string {
	symbols := "abcdefghijkl1234567890mnopqrstuvwxyz1234567890"

	b := make([]byte, n)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}
