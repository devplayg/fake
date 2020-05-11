package main

import (
	"fmt"
	"github.com/devplayg/fake"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(fake.Num(1, 10))
}
