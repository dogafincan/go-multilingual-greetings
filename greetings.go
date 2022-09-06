package greetings

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var greetings = []string{
	"안녕하세요",
	"こんにちは",
	"你好",
	"สวัสดีครับ",
	"Xin chào",
}

func RandomGreeting() string {
	return greetings[rand.Intn(len(greetings))]
}
