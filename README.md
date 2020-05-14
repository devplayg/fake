# fake

[![Build Status](https://travis-ci.org/devplayg/fake.svg?branch=master)](https://travis-ci.org/devplayg/fake)

Fake data generator in Go; Fake number, string, sleep, time, etc.

```go
fmt.Printf("Fake date: %s\n", fake.Now())
fmt.Printf("Fake username: %s\n", fake.LowerChars(fake.Num(6, 10)))
fmt.Printf("Fake IP: %s\n", fake.IPv4(192, 168, 0))
fmt.Printf("Fake credit-card: %s\n", fake.NumCode("4###-####-####-####"))
fmt.Printf("Fake phone number: %s\n", fake.NumCode("+01-555-###-####"))
fmt.Printf("Fake MAC address: %s\n", fake.Mac())
fmt.Printf("Fake OS: %s\n", fake.PickStr([]string{"Windows", "Linux", "Mac"}))
fmt.Printf("Fake age: %d\n", fake.Num(10, 20))
fmt.Printf("Fake favorite quot: %s\n", fake.Phrase())
fmt.Printf("Fake bool: %v\n", fake.Bool())
fmt.Println(strings.Repeat("-", 65))
```

Output

```
Fake date: 2020-05-14 16:46:54.3682151 +0900 KST m=-0.651090199
Fake username: edhujr
Fake IP: 192.168.0.68
Fake credit-card: 4783-4279-7305-1400
Fake phone number: +01-555-841-7622
Fake MAC address: f3:f3:a1:2f:b8:99
Fake OS: Linux
Fake age: 18
Fake favorite quot: dicta soluta tempore sapien
Fake bool: false
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
String(n int) string      // "a0B2c3", Alphanumeric  
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