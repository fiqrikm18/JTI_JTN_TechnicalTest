package util

import (
	"fmt"
	"math/rand"
	"time"
)

func GeneratePhoneNumber() string {
	rand.Seed(time.Now().UnixNano())
	phoneNumber := "08"

	for i := 1; i < 11; i++ {
		phoneNumber += fmt.Sprintf("%d", rand.Intn(10))
	}

	return phoneNumber
}
