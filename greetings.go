package greetings

import (
	"math/rand"
	"time"
)

var greetings = [5]string{
	"안녕하세요",
	"こんにちは",
	"你好",
	"สวัสดีครับ",
	"Xin chào",
}

func Hello() string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(greetings))
	return greetings[randomIndex]
}
