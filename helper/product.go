package helper

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateSKU() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 12

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var sb strings.Builder
	sb.WriteString("ITEM-")

	for range length {
		sb.WriteByte(charset[rng.Intn(len(charset))])
	}

	return sb.String()
}
