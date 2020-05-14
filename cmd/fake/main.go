package main

import (
	"context"
	"fmt"
	"github.com/devplayg/fake"
	"strings"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	for {
		fmt.Printf("Fake date: %s\n", fake.Now())
		fmt.Printf("Fake username: %s\n", fake.LowerChars(fake.Num(6, 10)))
		fmt.Printf("Fake IP: %s\n", fake.IPv4(192, 168, 0))
		fmt.Printf("Fake MAC address: %s\n", fake.Mac())
		fmt.Printf("Fake OS: %s\n", fake.PickStr([]string{"Windows", "Linux", "Mac"}))
		fmt.Printf("Fake speed: %d\n", fake.Num(10, 20))
		fmt.Printf("Fake phrase: %s\n", fake.Phrase())
		fmt.Printf("Fake bool: %v\n", fake.Bool())
		fmt.Println(strings.Repeat("-", 65))

		// fake.SleepMillis(100, 500)

		select {
		case <-ctx.Done():
			fmt.Println("good bye")
			return
		case <-fake.TimeMillisAfter(100, 500):
		}
	}
}
