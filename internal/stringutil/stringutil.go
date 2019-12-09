package stringutil

import (
	"math/rand"
	"time"
)

func CreateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	var str string
	for i := 0; i < length; i++ {
		val := rand.Intn(26)
		str += string(val + 'a')
	}
	return str
}
