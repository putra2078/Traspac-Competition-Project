package utils

import (
	"time"
	"math/rand"
	"fmt"
)

func GenerateNIPWithPrefix(prefix string) string {
	now := time.Now()
	datePart := now.Format("20060102")
	randomPart := rand.Intn(99999) + 1
	return fmt.Sprintf("%s-%s-%05d", prefix, datePart, randomPart)
}
