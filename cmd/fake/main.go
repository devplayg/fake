package main

import (
	"math/rand"
	"time"
)

const tryCount = 5

func main() {
	rand.Seed(time.Now().UnixNano())
	//
	//for i := 0; i < tryCount; i++ {
	//	fmt.Println(fake.Num(1, 10))
	//}
	//for i := 0; i < tryCount; i++ {
	//	fmt.Println(fake.NumStr(1, 10))
	//}
	//for i := 0; i < tryCount; i++ {
	//	fmt.Println(fake.Alpha(5))
	//}
	//for i := 0; i < tryCount; i++ {
	//	fmt.Println(fake.String(5))
	//}
	//
	//for i := 0; i < tryCount; i++ {
	//	fmt.Println(fake.IPv4())
	//}
	//for i := 0; i < tryCount; i++ {
	//	fmt.Println(fake.IPv4(192))
	//}
	//for i := 0; i < tryCount; i++ {
	//	fmt.Println(fake.IPv4(192, 168))
	//}
	//for i := 0; i < tryCount; i++ {
	//	fmt.Println(fake.IPv4(192, 168, 0, 11, 11))
	//}
	//for i := 0; i < tryCount; i++ {
	//	spew.Dump(fake.Byte(3))
	//}

	//mac,err := net.ParseMAC("00:00:5e:a00:53:01")
	//fmt.Println(mac.String())
	//spew.Dump(err)

	//t1,_ := time.Parse("2006-01-02 15:04:05", "2012-04-05 00:00:00")
	//for i := 0; i < tryCount; i++ {
	//fmt.Println(fake.DateRange(t1, t1.Add(3*time.Second)))
	//fmt.Println(fake.Date())
	//}

}
