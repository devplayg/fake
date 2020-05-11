# fake

[![Build Status](https://travis-ci.org/devplayg/fake.svg?branch=master)](https://travis-ci.org/devplayg/fake)

Fake data generator in Go

## 1. Normal works 

### Number

```go
// Numeric
Num(min, max int) int 
Num64(min, max int64) int64

// String
NumStr(min, max) string // "123"
String(n int) string // Alphanumeric "a0B2c3"  
Alpha(n int) string // "aBc"
UpperCase(n int) string // "ABC"
LowerCase(n int) string // "abc"
```

## 2. Security-sensitive works

### Number ()

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
``` 

### String


- [ ] DigitN(n int) string // "aaa"
- [ ] Char() string // "a"
- [ ] CharN() string "aBc"
- [ ] LowerCase() string // "a"
- [ ] LowerCaseN() string // "abc"
- [ ] UpperCase() string // "A"
- [ ] UpperCaseN() string // "ABC"

### Byte

- [ ] Byte() // size: 0..2147483647
- [ ] Byte(5) // size: 0..5
- [ ] Byte(10, 100) // size: 10..100

### Network

- [ ] IPv4() //
- [ ] IPv4(192) // A Class
- [ ] IPv4(192, 168) // B Class
- [ ] IPv4(192, 168, 0) // C Class
- [ ] IPv4But()


### Date

- [ ] Date

## Benchmark

    $ go test -bench=. -benchmem
    goos: linux
    goarch: amd64
    pkg: github.com/devplayg/fake
    BenchmarkInt8-2       	48413200	        23.9 ns/op	       0 B/op	       0 allocs/op
    BenchmarkInt16-2      	50145693	        23.8 ns/op	       0 B/op	       0 allocs/op
    BenchmarkInt32-2      	50781237	        24.2 ns/op	       0 B/op	       0 allocs/op
    BenchmarkInt64-2      	 1000000	      1046 ns/op	     168 B/op	       8 allocs/op
    BenchmarkUInt8-2      	50676717	        23.6 ns/op	       0 B/op	       0 allocs/op
    BenchmarkUInt16-2     	50824243	        23.5 ns/op	       0 B/op	       0 allocs/op
    BenchmarkUInt32-2     	50886445	        23.5 ns/op	       0 B/op	       0 allocs/op
    BenchmarkUInt64-2     	 1252040	       979 ns/op	     120 B/op	       7 allocs/op
    BenchmarkFloat32-2    	53117192	        22.6 ns/op	       0 B/op	       0 allocs/op
    BenchmarkFloat642-2   	57842583	        20.7 ns/op	       0 B/op	       0 allocs/op
    PASS
    ok  	github.com/devplayg/fake	13.027s