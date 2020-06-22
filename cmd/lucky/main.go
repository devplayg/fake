package main

import (
	"fmt"
	"github.com/devplayg/fake"
	"os"
	"strings"
	"time"
)

const (
	Count = 6
	Min   = 1
	Max   = 45
)

func main() {
	displayLogo("Aloha numbers")
	balls := createBalls(Min, Max)
	var num int
	line := fmt.Sprintf("|%s|", time.Now().Format(time.RFC3339))
	for i := 0; i < Count; i++ {
		fake.ShuffleInts(balls)
		num, balls = balls[0], balls[1:]
		line += fmt.Sprintf("%d|", num)
	}
	write(strings.TrimSpace(line))
}

func createBalls(min, max int) []int {
	balls := make([]int, max, max)
	for i := min; i <= max; i++ {
		balls[i-1] = i
	}
	return balls
}

func write(line string) {
	file, err := os.OpenFile("README.md", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.WriteString(line + "\n"); err != nil {
		panic(err)
	}
}

func displayLogo(title string) {
	loading := "/-\\|"
	arr := strings.Split(title, "")
	var line string

	for k := 0; k < Count; k++ {
		line = ""
		fmt.Printf("%s\r", strings.Repeat(" ", len(title)+5))
		for i := 0; i < len(title); i++ {
			line += arr[i]
			fmt.Printf("[%c] %s\r", loading[i%4], line)
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Printf("[V] %s: completed\n", line)
}
