package greetings

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var greetings = [5]string{
	"안녕하세요",
	"こんにちは",
	"你好",
	"สวัสดีครับ",
	"Xin chào",
}

func Hello() string {
	randomIndex := rand.Intn(len(greetings))
	return greetings[randomIndex]
}
