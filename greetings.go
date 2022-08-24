package greetings

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomGreeting() string {
	greetings := [5]string{
		"안녕하세요",
		"こんにちは",
		"你好",
		"สวัสดีครับ",
		"Xin chào",
	}

	return greetings[rand.Intn(len(greetings))]
}
