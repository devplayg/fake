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
		fmt.Printf("Fake USERNAME: %s\n", fake.UpperChars(fake.Num(6, 10)))
		fmt.Printf("Fake IP: %s\n", fake.IPv4(192, 168, 0))
		fmt.Printf("Fake credit-card: %s\n", fake.NumCode("4###-####-####-####"))
		fmt.Printf("Fake phone number: %s\n", fake.NumCode("+01-555-###-####"))
		fmt.Printf("Fake MAC address: %s\n", fake.Mac())
		fmt.Printf("Fake OS: %s\n", fake.PickStr([]string{"Windows", "Linux", "Mac"}))
		fmt.Printf("Fake age: %d\n", fake.Num(10, 20))
		fmt.Printf("Fake favorite quot: %s\n", fake.Phrase())
		fmt.Printf("Fake bool: %v\n", fake.Bool())
		fmt.Printf("Fake UUID: %s\n", fake.UUID())
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
