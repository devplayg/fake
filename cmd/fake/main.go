package main

import (
	"fmt"
	"github.com/devplayg/fake"
	"math/rand"
	"time"
)

const tryCount = 5

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < tryCount; i++ {
		fmt.Println(fake.Num(1, 10))
	}
	for i := 0; i < tryCount; i++ {
		fmt.Println(fake.NumStr(1, 10))
	}
	for i := 0; i < tryCount; i++ {
		fmt.Println(fake.Alpha(5))
	}
	for i := 0; i < tryCount; i++ {
		fmt.Println(fake.String(5))
	}

	for i := 0; i < tryCount; i++ {
		fmt.Println(fake.IPv4())
	}
	for i := 0; i < tryCount; i++ {
		fmt.Println(fake.IPv4(192))
	}
	for i := 0; i < tryCount; i++ {
		fmt.Println(fake.IPv4(192, 168))
	}
	for i := 0; i < tryCount; i++ {
		fmt.Println(fake.IPv4(192, 168, 0, 11, 11))
	}
}
func aaa(f func(n int) string) {
	f(5)
}
