# fake

[![Build Status](https://travis-ci.org/devplayg/fake.svg?branch=master)](https://travis-ci.org/devplayg/fake)
[![Go Report Card](https://goreportcard.com/badge/github.com/devplayg/fake)](https://goreportcard.com/report/github.com/devplayg/fake)

Fake data generator in Go; Fake number, string, sleep, time, etc.

```go
fake.Bool()                                 // false
fake.Digit()                                // "3"
fake.Num(10, 20)                            // 18                                                                       
fake.UpperChar()                            // "A"
fake.UpperChars(5)                          // ALOHA
fake.LowerChar()                            // "a"                                                                   
fake.LowerChars(5)                          // "aloha"
fake.PickStr([]string{"Windows", "Linux"})  // "Linux"                                                                    
fake.IPv4(192, 168, 0))                     // "192.168.0.68"                                                             
fake.NumCode("4###-####-####-####")         // "4783-4279-7305-1400" (Credit-card)                                                      
fake.NumCode("+01-555-###-####")            // "+01-555-841-7622" (Phone)                                                         
fake.Mac()                                  // "f3:f3:a1:2f:b8:99"                                                        
fake.Now()                                  // "2020-05-14 16:46:54.3682151 +0900 KST m=-0.651090199"                     
fake.Phrase()                               // "dicta soluta tempore sapien" (Lorem Ipsum)                                              
fake.UUID()                                 // "5cab345e-1f78-4d07-9622-53f0a1d057fb"
```

## 1. Normal works 

### Numeric

```go
Num(min, max int) int 
Num64(min, max int64) int64
NumStr(min, max int) string // "123"
Digit() string // "1"
PicInt([]int) int
```

### String

```go
String(n int) string       // "a0B2c3", Alphanumeric  
Alpha(n int) string        // "aBc"
LowerChar() string         // "a"
LowerChars(n int) string   // "abc"
UpperChar() string         // "A"
UpperChars(n int) string   // "ABC"
PicStr([]string) string    // one of them
NumCode(str string) string // "555-###-####", "4###-####-####-####"
```

### Network

```go
IPv4() string               // "10.10.10.10"
IPv4(int) string            // "192.10.10.10"
IPv4(int, int) string       // "192.168.10.10"
IPv4(int, int, int) string  // "192.168.54.10"
IPv6() string               // "2001:0000:a423:5690:6d9f:b71d:b488:f758"
Mac() string                // "00:00:5e:00:53:01"
MacHalf() string            // "00:53:01"
```

### Byte

```go
Byte(n int) []byte
```

### Date & time

```go
Now() time.Time             // ± 3000 Milliseconds
Date(t time.Time) time.Time // ± 3000 Milliseconds
DateWithJitter(t time.Time, jitter time.Duration) time.Time // ± jitter
DateRange(min, max time.Time) time.Time
SleepMillis(min, max uint32)
TimeMillisAfter(min, max uint32) <-chan time.Time
```

### Image

```go
Jpeg(width int, height int) []byte
Png(width int, height int) []byte
Gif(width int, height int) []byte
Image(width, height int) *image.RGBA
```

### Etc.

```go
Hex() string // "ff"
```

### Lorem ipsum

```go
Word() string
Phrase() string
Sentence() string
Paragraph() string
```

### Shuffling

```go
ShuffleInts(arr []int)
ShuffleStr(arr []string)
```

## 2. Sensitive works

### Number

```go
Int8() int8 
Int16() int16 
Int32() int32 
Int64() int64 
Int8Range(min, max int8) int8 
Int16Range(min, max int16) int16 
Int32Range(min, max int32) int32 
Int64Range(min, max int64) int64 
Uint8() uint8 
Uint16() uint16 
Uint32() uint32 
Uint64() uint64 
Uint8Range(min, max uint8) uint8 
Uint16Range(min, max uint16) uint16 
Uint32Range(min, max uint32) uint32 
Uint64Range(min, max uint64) uint64 
Float32() float32 
Float32Range(min, max float32) float32 
Float64() float64 
Float64Range(min, max float64) float64
BigIntRange(min, max *big.Int) *big.Int // Big int
``` 

## Benchmark

    $ go test -bench=. -benchmem
    goos: linux
    goarch: amd64
    pkg: github.com/devplayg/fake
    BenchmarkInt8-2       	48413200	        23.9 ns/op	       0 B/op	       0 allocs/op
    BenchmarkInt16-2      	50145693	        23.8 ns/op	       0 B/op	       0 allocs/op
    BenchmarkInt32-2      	50781237	        24.2 ns/op	       0 B/op	       0 allocs/op
    BenchmarkInt64-2      	 1000000	        1046 ns/op	     168 B/op	       8 allocs/op
    BenchmarkUInt8-2      	50676717	        23.6 ns/op	       0 B/op	       0 allocs/op
    BenchmarkUInt16-2     	50824243	        23.5 ns/op	       0 B/op	       0 allocs/op
    BenchmarkUInt32-2     	50886445	        23.5 ns/op	       0 B/op	       0 allocs/op
    BenchmarkUInt64-2     	 1252040	        979 ns/op	     120 B/op	       7 allocs/op
    BenchmarkFloat32-2    	53117192	        22.6 ns/op	       0 B/op	       0 allocs/op
    BenchmarkFloat642-2   	57842583	        20.7 ns/op	       0 B/op	       0 allocs/op
    PASS
    ok  	github.com/devplayg/fake	13.027s