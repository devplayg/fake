package main

import (
	"fmt"
	"github.com/devplayg/fake"
	mrand "math/rand"
	"time"
)

func main() {

	//spew.Dump(fake.Int64())
	//spew.Dump(fake.Int64())
	//spew.Dump(fake.Int64())
	//spew.Dump(fake.Int64())
	//spew.Dump(fake.Int64())

	//const MaxUint = ^uint(0)
	mrand.Seed(time.Now().UnixNano())
	//
	//n1 :=big.NewInt(int64(math.MinInt64))
	//n2 := big.NewInt(int64(math.MaxInt64))
	//c := big.NewInt(0)
	//n := c.Sub(n2, n1).Add(c,big.NewInt(1))
	//randNum, _ := crand.Int(crand.Reader, n)
	//randNum.Add(randNum, n1)
	//fmt.Println(randNum.Int64())

	//fmt.Println( int(MaxUint >> 1)  )
	//fmt.Println( math.MinInt64 )
	//fmt.Println( math.MaxInt64 )
	for {
		//fmt.Println(fake.Number(-3, 3))
		//a := mrand.Intn(26) + 'a'
		// spew.Dump(fake.Num(math.MinInt64,math.MaxInt64))
		// min := big.NewInt(math.MinInt64)
		// max := big.NewInt(math.MaxInt64)
		// c := new(big.Int)
		// c = c.Sub(max, min)

		var n int
		n = 123412384712983749812374
		fmt.Println(n)
		fmt.Println(fake.Num64(-3, -1))

		//fmt.Println(fake.NumberWithBig(math.MinInt64, math.MaxInt64))
		time.Sleep(100 * time.Millisecond)
	}
	//fmt.Println(fake.NumberWithBig(math.MinInt32, math.MaxInt64))
	//fmt.Println(fake.Bool())

	//var max int32 = math.MaxInt32 // 16
	//var min int32  = math.MinInt32 // -15
	//var n int32 // 31
	//n = max - min
	//fmt.Println(n)
	//
	//r := rand.Intn(-1)
	//fmt.Println(r)

	//fmt.Println(fake.Number(-10, -5))
	//fmt.Println(fake.Int64Range(-10, -1))
	//fmt.Println(fake.Int64())
	//fmt.Println(math.MinInt64)
	//fmt.Println(fake.Int32())

	//fmt.Println(math.MinInt32)
	//fmt.Println(fake.NumberWithBig(-2, 2))

}
