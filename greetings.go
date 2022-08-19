package greetings

import (
	"fmt"
	"math/rand"
	"time"
)

func Hello() {
	greetings := [5]string{
		"안녕하세요",
		"こんにちは",
		"你好",
		"สวัสดีครับ",
		"Xin chào",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(greetings))
	fmt.Println(greetings[randomIndex])
}
