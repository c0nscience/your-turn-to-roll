package password

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"strings"
)

var (
	alpha   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numeric = "0123456789"

	combined = alpha + numeric
)

func Generate(length int) string {
	combinedRunes := []rune(combined)
	rand.Shuffle(len(combinedRunes), func(i, j int) {
		combinedRunes[i], combinedRunes[j] = combinedRunes[j], combinedRunes[i]
	})

	combined = string(combinedRunes)

	password := strings.Builder{}

	for i := 0; i < length; i++ {
		n, _ := crand.Int(crand.Reader, big.NewInt(int64(len(combined))))
		password.WriteString(string(combined[n.Int64()]))
	}

	runes := []rune(password.String())
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
